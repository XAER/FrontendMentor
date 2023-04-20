// import the fs module
import { readdirSync } from "fs";

const challengesGrid = document.getElementById("challenges-grid");

// get a list of the folders in this folder
const challenges = readdirSync(__dirname);

// loop through the folders
challenges.forEach((challenge) => {
  // create a new div element
  const challengeDiv = document.createElement("div");
  // add the challenge class to the div
  challengeDiv.classList.add("challenge");
  // create a new link element
  const challengeLink = document.createElement("a");
  // set the href of the link to the challenge folder
  challengeLink.href = `./${challenge}/index.html`;
  // set the text of the link to the challenge folder
  challengeLink.textContent = challenge;
  // append the link to the div
  challengeDiv.appendChild(challengeLink);
  // append the div to the challenges grid
  challengesGrid.appendChild(challengeDiv);
});
