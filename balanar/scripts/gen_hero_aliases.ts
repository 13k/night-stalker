import { promises as fs } from "fs";

import _ from "lodash";
import vdf from "simple-vdf";

const usage = (): never => {
  console.log("Usage: %s <npc_heroes.txt> <output.json>", process.argv[1]);
  process.exit(1);
};

interface HeroesAliases {
  [heroName: string]: string[];
}

const transformKeyValues = (kv: vdf.KeyValues): HeroesAliases =>
  _.chain(kv.DOTAHeroes as vdf.KeyValues)
    .pickBy(
      (hero: vdf.KeyValues, name: string) =>
        name.match(/^npc_dota_hero_/) != null && hero.Enabled === "1"
    )
    .transform((heroes: HeroesAliases, hero: vdf.KeyValues, name: string) => {
      const nameAliases: string = hero.NameAliases as string;

      if (_.isEmpty(nameAliases)) {
        return;
      }

      heroes[name] = _.chain(nameAliases)
        .split(/[,;]/)
        .map(_.trim)
        .map(_.toLower)
        .sortBy()
        .value();
    }, {})
    .value();

const save = (aliases: HeroesAliases, output: string): Promise<void> =>
  fs.writeFile(output, JSON.stringify(aliases, null, 2), { encoding: "utf8" });

const main = (): Promise<void> => {
  if (process.argv.length < 4) {
    usage();
  }

  const [, , input, output] = process.argv;

  return fs
    .readFile(input, { encoding: "utf8" })
    .then(vdf.parse)
    .then(transformKeyValues)
    .then(aliases => save(aliases, output))
    .then(() => {
      console.log("File saved to %s", output);
    })
    .catch(console.error);
};

main();
