import { promises as fs, createWriteStream, Stats } from "fs";

import _ from "lodash";
import vdf from "simple-vdf";
import genericPool from "generic-pool";
import HttpAgent, { HttpsAgent } from "agentkeepalive";
import { $enum } from "ts-enum-util";
import got, {
  Got,
  ExtendOptions as GotExtendOptions,
  Headers as GotHeaders,
  Response as GotResponse,
} from "got";

const usage = (): never => {
  console.log("Usage: %s <npc_heroes.txt> <output_dir>", process.argv[1]);
  process.exit(1);
};

enum Version {
  full,
  vert,
}

const $Version = $enum(Version);

const VERSION_EXT = {
  [Version.full]: "png",
  [Version.vert]: "jpg",
};

const CDN_URL = "http://cdn.dota2.com";

const HTTP_CLIENT_OPTIONS: GotExtendOptions = {
  prefixUrl: CDN_URL,
  decompress: true,
  responseType: "buffer",
  retry: 0,
  throwHttpErrors: true,
  agent: {
    http: new HttpAgent(),
    https: new HttpsAgent(),
  },
};

class Asset {
  heroName: string;
  shortHeroName: string;
  version: Version;
  versionName: string;
  remoteFilename: string;
  remotePath: string;
  directory: string;
  filename: string;
  ext: string;
  path: string;
  stats: Stats;

  constructor(heroName: string, version: Version, directory: string) {
    this.heroName = heroName;
    this.shortHeroName = heroName.replace(/^npc_dota_hero_/, "");
    this.version = version;
    this.versionName = Version[this.version];
    this.directory = directory;
    this.ext = VERSION_EXT[this.version];
    this.remoteFilename = `${this.shortHeroName}_${this.versionName}.${this.ext}`;
    this.remotePath = `apps/dota2/images/heroes/${this.remoteFilename}`;
    this.filename = `${this.heroName}_${this.versionName}.${this.ext}`;
    this.path = `${this.directory}/${this.filename}`;
  }

  stat(): Promise<Stats | null> {
    return fs
      .stat(this.path)
      .then(stats => {
        this.stats = stats;
        return this.stats;
      })
      .catch(err => {
        if (err.code === "ENOENT") {
          this.stats = null;
          return this.stats;
        }

        throw err;
      });
  }
}

enum DownloadFileResult {
  Created,
  Updated,
  NotModified,
}

const downloadFileResultSymbol = (result: DownloadFileResult): string => {
  return $enum.mapValue(result).with({
    [DownloadFileResult.Created]: "+",
    [DownloadFileResult.Updated]: "!",
    [DownloadFileResult.NotModified]: "=",
  });
};

interface DownloadResult {
  asset: Asset;
  fileResult: DownloadFileResult;
}

const transformKeyValues = (kv: vdf.KeyValues): string[] =>
  _.chain(kv.DOTAHeroes as vdf.KeyValues)
    .pickBy(
      (hero: vdf.KeyValues, name: string) =>
        name.match(/^npc_dota_hero_/) != null && hero.Enabled === "1"
    )
    .keys()
    .sortBy()
    .value();

const download = (conn: Got, asset: Asset): Promise<DownloadResult> => {
  return asset
    .stat()
    .then(stats => {
      const headers: GotHeaders = {};

      if (stats != null) {
        headers["if-modified-since"] = stats.mtime.toUTCString();
      }

      return headers;
    })
    .then(headers => conn.get(asset.remotePath, { headers, responseType: "buffer" }))
    .then(res => handleResponse(res, asset));
};

const handleResponse = (res: GotResponse<Buffer>, asset: Asset): Promise<DownloadResult> =>
  new Promise((resolve, reject) => {
    switch (res.statusCode) {
      case 200:
        const outputStream = createWriteStream(asset.path);

        outputStream.on("error", reject);

        outputStream.end(res.body, () => {
          const fileResult =
            asset.stats != null ? DownloadFileResult.Updated : DownloadFileResult.Created;

          resolve({ fileResult, asset });
        });

        break;
      case 304:
        resolve({ fileResult: DownloadFileResult.NotModified, asset });
        break;
      default:
        reject(
          new Error(
            `${asset.heroName} [${asset.versionName}]: invalid response code ${res.statusCode}`
          )
        );
    }
  });

const queueDownload = (pool: genericPool.Pool<Got>, asset: Asset): Promise<DownloadResult> =>
  (pool.acquire() as Promise<Got>).then(conn =>
    download(conn, asset)
      .then(result => {
        console.log("[%s] %s", downloadFileResultSymbol(result.fileResult), result.asset.path);
        return result;
      })
      .finally(() => {
        pool.release(conn);
      })
  );

const downloadAll = (
  pool: genericPool.Pool<Got>,
  names: string[],
  outputDir: string
): Promise<DownloadResult>[] =>
  _.chain($Version.getValues())
    .map(version => _.map(names, name => queueDownload(pool, new Asset(name, version, outputDir))))
    .flatten()
    .value();

const main = (): Promise<void> => {
  if (process.argv.length < 4) {
    usage();
  }

  const [, , input, outputDir] = process.argv;

  const poolFactory: genericPool.Factory<Got> = {
    create: () => Promise.resolve(got.extend(HTTP_CLIENT_OPTIONS)),
    destroy: (conn: Got) => Promise.resolve(),
  };

  const pool = genericPool.createPool(poolFactory, { min: 1, max: 10 });

  return fs
    .readFile(input, { encoding: "utf8" })
    .then(vdf.parse)
    .then(transformKeyValues)
    .then(names => downloadAll(pool, names, outputDir))
    .then(downloads => Promise.allSettled(downloads))
    .then(results => {
      const failed = _.filter(results, { status: "rejected" });

      _.each(failed, (result: PromiseRejectedResult) => {
        const err = result.reason;

        if (err instanceof got.HTTPError) {
          const { url } = err.options;
          console.error("%s: %s", url.href, err.message);
        } else {
          console.error(err);
        }
      });
    })
    .catch(console.error);
};

main();
