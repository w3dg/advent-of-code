import { removeEmptyElement } from "./utils";

export function terminalOutputParser(terminalOutput) {
  const commandAndOutput = removeEmptyElement(terminalOutput.split("$ "));
  const commandAndOutputArr = commandAndOutput.map((co) => co.split("\n"));

  const commandAndOutputMapped = commandAndOutputArr.map((x) => removeEmptyElement(x));

  const fileStructure = {};

  commandAndOutputMapped.forEach((commandAndOutput) => {
    if (commandAndOutput.startsWith("cd")) {
      cd();
    } else if (commandAndOutput[0] === "ls") {
      lsOutput = commandAndOutput.splice(0);
      ls(lsOutput);
    }
  });
}
