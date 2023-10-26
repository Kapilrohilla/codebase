const fs = require("fs/promises");

(async () => {
  const createFile = async (path) => {
    try {
      // checking if file already exists
      const exisitingFileHandle = await fs.open(path, "r");
      // file already exists
      exisitingFileHandle.close();

      return console.log(`Create: file ${path} already exists.`);
    } catch (err) {
      // File not exists, Thus creating a file
      const newFileHandle = await fs.open(path, "w");
      // file created now close successfully..
      await newFileHandle.close();
      console.log(`Create: Successfully created! ${path}`);
    }
  };

  const deleteFile = async (path) => {
    try {
      await fs.unlink(path);
      console.log(`Delete :${path} deleted successfully`);
    } catch (err) {
      if ((err.code = "ENOENT")) {
        console.log(`Delete: ${path} not exists`);
      } else {
        console.log(err);
      }
    }
  };

  const renameFile = async (oldPath, newPath) => {
    try {
      const exisitingFileHandler = await fs.open(oldPath, "r");
      await exisitingFileHandler.close();

      await fs.rename(oldPath, newPath);
      console.log(`Rename: "${oldPath}" renamed to "${newPath}"`);
    } catch (err) {
      if ((err.code = "ENOENT")) {
        console.log(`Rename: ${oldPath} not exists`);
      } else {
        console.log(err);
      }
    }
  };
  const addToFile = async (path, data) => {
    try {
      const exisitingFileHandler = await fs.open(path, "r");
      await exisitingFileHandler.close();

      // const encoding = "utf-8";
      await fs.appendFile(path, data);
      console.log(`AddTo: "${data}" appended successfully, to ${path}`);
    } catch (err) {
      if ((err.code = "ENOENT")) {
        console.log(`AddTo : ${path} not exists`);
      } else {
        console.log(err);
      }
    }
  };
  // COMMANDS
  const CREATE_FILE = "create a file";
  const DELETE_FILE = "delete a file";
  const RENAME_FILE = "rename a file";
  const AddTO_FILE = "add to file";

  const commandFileHandler = await fs.open("./command.txt", "r");

  commandFileHandler.on("change", async () => {
    /// get the size of our file
    const { size } = await commandFileHandler.stat();
    // allocate our buffer with same size as of the file
    const buff = Buffer.alloc(size);
    // location at which we want to read the buffer
    const offset = 0;
    // how many bytes  we want to read
    const length = buff.byteLength;
    // location at which we want to read the file.
    const position = 0;
    // whole data after reading the file
    await commandFileHandler.read(buff, offset, length, position);
    const command = buff.toString("utf-8");
    // CREATE A FILE
    /// create a file <path>
    if (command.includes(CREATE_FILE)) {
      const filepath = command.substring(CREATE_FILE.length + 1).trim(); // here 14 = 'create a file ' + 1

      createFile(filepath);
    }
    if (command.includes(DELETE_FILE)) {
      const filePath = command.substring(DELETE_FILE.length + 1);

      deleteFile(filePath);
    }

    if (command.includes(RENAME_FILE)) {
      const idx = command.indexOf(" to ");
      const oldFilePath = command.substring(RENAME_FILE.length + 1, idx);
      const newFilePath = command.substring(idx + 4);

      renameFile(oldFilePath, newFilePath);
    }

    if (command.includes(AddTO_FILE)) {
      const idx = command.indexOf(" this content: ");

      if (idx === -1) {
        return console.log('AddTo: Syntax Error, use " this content: "');
      }
      const filePath = command.substring(AddTO_FILE.length + 1, idx).trim();
      console.log(filePath, 0);
      const data = command.substring(idx + " this content: ".length).trim();
      console.log(data, 1);

      addToFile(filePath, data);
    }
  });

  const watcher = fs.watch("./command.txt");

  for await (const event of watcher) {
    if (event.eventType === "change") {
      commandFileHandler.emit("change");

      //
    }
    // commandFileHandler.close();
  }
})();
