/* eslint-disable no-console */

const util = require("util");
const fs = require("fs");
const readFile = util.promisify(fs.readFile);
const writeFile = util.promisify(fs.writeFile);

const _ = require("lodash");
const vdf = require("simple-vdf");

function usage() {
  console.log("Usage: %s <npc_heroes.txt> <output.json>", process.argv[1]);
  process.exit(1);
}

function generate(kv) {
  return _.chain(kv.DOTAHeroes)
    .toPairs()
    .filter(([name]) => {
      return name !== "npc_dota_hero_base" && name.match(/^npc_dota_hero_/);
    })
    .map(([name, heroKV]) => {
      if (_.isEmpty(heroKV.NameAliases)) {
        return [name, []];
      }

      const aliases = _.chain(heroKV.NameAliases)
        .split(/[,;]/)
        .map(_.trim)
        .map(_.toLower)
        .sortBy()
        .value();

      return [name, aliases];
    })
    .fromPairs()
    .value();
}

function save(aliases, output) {
  const encoded = JSON.stringify(aliases, null, 2);
  return writeFile(output, encoded, { encoding: "utf8" });
}

if (process.argv.length < 4) {
  usage();
}

const [, , input, output] = process.argv;

readFile(input, { encoding: "utf8" })
  .then(vdf.parse)
  .then(generate)
  .then(aliases => save(aliases, output))
  .then(() => {
    console.log("File saved to %s", output);
  })
  .catch(err => console.error(err));
