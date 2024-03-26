#!/usr/bin/env node

const fs = require("fs");
const path = require("path");
const readline = require("readline/promises");
const fileName = "todo.log";

const logPath = path.join(__dirname, fileName);
const completedTaskFile = path.join(__dirname, "completed.log");
const argv = process.argv;

const hashmap = {
  add: add_todo,
  list: list_todo,
  comp: comp_todo,
  del: del_todo,
  "list:comp": list_comp,
};

function main() {
  if (argv.length < 3) {
    return help();
  }
  const hk = Object.keys(hashmap);

  let found = false;
  hk.forEach((key) => {
    if (key === argv[2]) {
      hashmap[key]();
      found = true;
    }
  });

  if (!found) {
    console.error("Invalid command");
    return;
  }
}

main();

function help() {
  console.log(`============== Welcome to TODO_CLI ==============
  usage <command>
  commands can be:
  add:  create todo
  comp: complete todo
  del:  delete todo
  list: list all pending todo
  list:comp: list all completed todos
  `);
}
/// no change require
function add_todo() {
  if (!argv[3]) {
    return console.error("Error: text not found to add");
  }
  fs.appendFile(logPath, `${argv[3]}\n`, (err) => {
    if (err) {
      return console.error(err);
    } else {
      console.log("Added successfully.");
    }
  });
}
// no change require
function list_todo() {
  fs.readFile(logPath, "utf-8", (err, data) => {
    if (err) {
      console.error("File not found");
      fs.writeFile(logPath, "", (err, data) => {
        if (err) {
          console.error(err);
        } else {
        }
      });
    } else {
      const todoArr = data.trim().split("\n");
      let lineCount = 1;
      console.log(`********** TODO COUNT = ${todoArr.length} **********`);
      todoArr.map((todo) => {
        console.log(`${lineCount}. ${todo}`);
        lineCount++;
      });
    }
  });
}

async function comp_todo() {
  if (!argv[3]) return console.log("Error, please provide Task number");
  const isFileExists = fs.existsSync(completedTaskFile);
  if (!isFileExists) {
    fs.open(completedTaskFile, "w", (err, fd) => {
      if (err) {
        console.error("Error: " + err, 0);
      }
    });
  } else {
  }
  const rstream = fs.createReadStream(logPath);
  const rline = new readline.Interface({ input: rstream });

  let lineCounter = 1;
  for await (let line of rline) {
    if (lineCounter === Number(argv[3])) {
      fs.appendFile(completedTaskFile, line + "\n", (err) => {
        if (err) return console.log("Error: " + err);
      });
      break;
    }
    lineCounter++;
  }

  // delete line from todo
  deleteArgv3Line();
}
// deletion
function del_todo() {
  if (!argv[3]) {
    return console.log("Error, please provide Task number");
  }
  deleteArgv3Line();
}

function deleteArgv3Line() {
  const newFilePath = path.join(__dirname, "todo_temp.log");
  fs.writeFile(newFilePath, "", async (err) => {
    if (err) return console.error("failed to create new file");
    const rstream = fs.createReadStream(logPath);
    const readingLineByLine = readline.createInterface({ input: rstream });
    let line_count = 1;
    for await (const line of readingLineByLine) {
      if (line_count !== Number(argv[3])) {
        fs.appendFile(newFilePath, `${line}\n`, (err, data) => {
          if (err) console.error("Error: " + err);
        });
      }
      line_count++;
    }
    fs.unlink(logPath, (err) => {
      if (err) {
        return console.log(err, 1);
      }
      fs.rename(newFilePath, logPath, (err) => {
        if (err) {
          return console.log("Errror: " + err, 1);
        }
      });
    });
  });
}

function list_comp() {
  fs.readFile(completedTaskFile, "utf-8", (err, data) => {
    if (err) {
      console.error("File not found");
      fs.writeFile(logPath, "", (err, data) => {
        if (err) {
          console.error(err);
        } else {
        }
      });
    } else {
      const todoArr = data.trim().split("\n");
      let lineCount = 1;
      console.log(`********** TODO COUNT = ${todoArr.length} **********`);
      todoArr.map((todo) => {
        console.log(`${lineCount}. ${todo}`);
        lineCount++;
      });
    }
  });
}
