#! /usr/bin/env node

const yargs = require("yargs");
const utils = require("./utils");

const translate = require("[@vitalets/google-translate-api](http://twitter.com/vitalets/google-translate-api)");

if (sentence == "") {
  console.error("\nThe entered sentece is like John Cena, I can't see it\n");
  console.log("Enter tran --help to get started.\n");
}
translate(sentence, { to: language })
  .then((res) => {
    console.log("\n" + "\n" + res.text + "\n" + "\n");
  })
  .catch((err) => {
    console.error(err);
  });
// })

const usage = "\nUsage: tran <lang_name> sentence to be translated.";
const options = yargs
  .usage(usage)
  .option("l", {
    alias: "languages",
    describe: "List all supported languages.",
    type: "boolean",
    demandOption: false,
  })
  .help(true).argv;

console.log(yargs.argv._, 0);
console.log(process.argv, 1);

console.log("Hello world", 2);

let sentence = utils.parseSentence(yargs.argv._);
console.log(sentence, 3);
if (yargs.argv._[0] == null) {
  utils.showHelp;
  return;
}
if (yargs.argv.l == true || yargs.argv.languages == true) {
  utils.showAll();
  return;
}
if (yargs.argv._[0]) {
  var language = yargs.argv._[0].toLowerCase(); // stores the language
  //  parsing the language specified to the ISO-169-1 code.
  language = utils.parseLanguage(language);
}
