<p align="center">
  <a href="https://github.com/EnAccess/OpenSmartMeter">
    <img
      src="https://drive.google.com/uc?id=1gtL_p7l3HbOcCzc09A7KW5d7B5qn-BDs"
      alt="OpenSmartMeter"
      width="640"
    >
  </a>
</p>
<p align="center">
    9-10 May | Open Source in Energy Access Symposium Hackathon
</p>

---

## Port OpenPAYGO Python library to other languages (e.g. JS, PHP)

**Stack:** Python, either JS or PHP

**Helpful experiences:** Library programming and management, CI/CD pipelines

**Abstract:** The OpenPAYGO suite currently provides only a Python library.
Libraries in other programming languages whould increase and enhance the adaptability of the OpenPAYGO ecosystem. As such, they should be published to common, language-specific package repositories (for example, NPM for JS, packagist for PHP, etc…)

## Challenge

The goal of this challenge is to make the OpenPAYGO functionality available in other programming languages. JS and PHP have been discussed and raised in the community in the past, but if the participants see a need for other relevant languages, this is great as well.

_Note:_ The OpenPAYGO library encompasses different features of the OpenPAYGO ecosystem, like Token or Metrics. The focus of the workgroup should be the Token. Adding other OpenPAYGO features could be subject to future improvements outside of this hackathon.

> [!NOTE]
> In this repository we are focusing on JavaScript implementation
> It is perfectly fine to also port it to other languages if participants bring expertise.
> Repositories will be created on the fly in this case.

## Expected outcome

Like other challenges, the outcome of this one depends on the actual priorities and skill sets brought to the team.
A minimum expected outcome, however, is: One draft of a library in one additional relevant language, a roadmap, and a detailed to-do list of steps required to get to a final and usable level.

A non-exclusive list would look like this

1. Extracting test cases for all Token modes from the Python library (https://github.com/EnAccess/OpenPAYGO-Token/tree/main/tests)
2. Implement the OpenPAYGO Token algorithm in JavaScript
3. Add test cases based on the result of 1.
4. Integrate CI/CD pipeline to release the library to NPM

**Bonus outcome:** If the chosen language is JavaScript: Create a small static (!) website which can be used to generate OpenPAYGO Tokens for a testing use-case. The result should be similar to the example from Victron: https://payg.victronenergy.com/ (however Tokens should be computed client side, not server side like in the Victron case).

## Getting Started

- Join the OSEAS24 Discord server: https://community.oseas.org/
- Introduce yourself in #introductions channel and join this topic’s channel
- Confirm you have access to the following Repos
  - https://github.com/EnAccess/OpenPAYGO-js
  - https://github.com/EnAccess/OpenPAYGO-php
- For physical participants: Bring a computer (and required Adapters) for some hacking 🤖🧑‍💻
- Read the documentation
  - https://enaccess.github.io/OpenPAYGO-docs/
  - https://www.paygops.com/openpaygotoken

Contact person(s): Vivien / Daniel

## Further information and resources

- https://github.com/EnAccess/OpenPAYGO-python/issues/11
- https://github.com/EnAccess/OpenPAYGO-python/issues/12
