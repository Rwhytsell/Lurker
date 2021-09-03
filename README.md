[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<p>
  <a href="https://github.com/Rwhytsell/Lurker/">
    <!-- <img src="images/logo.png" alt="Logo" width="80" height="80"> -->
  </a>

  <h3 align="center">Lurker</h3>

  <div align="center">
    A twitch chat scraper
    <br />
    <!-- <a href="https://github.com/Rwhytsell/Lurker/"><strong>Explore the docs »</strong></a>
    <br /> -->
    <br />
    <a href="https://github.com/Rwhytsell/Lurker/">View Demo</a>
    ·
    <a href="https://github.com/Rwhytsell/Lurker/issues">Report Bug</a>
    ·
    <a href="https://github.com/Rwhytsell/Lurker/issues">Request Feature</a>
  </div>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

### Built With

* [Golang](https://golang.org/)
* [RabbitMQ](https://www.rabbitmq.com/)

<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* [Follow instructions to install Golang](https://go.dev/learn/)
* [Follow instructions to install RabbitMQ](https://www.rabbitmq.com/download.html)

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/Rwhytsell/Lurker/.git
   ```
2. Install Go packages
   ```sh
   go install
   ```
3. Build the project
      ```sh
      go build
      ```



<!-- USAGE EXAMPLES -->
## Usage

Lurker uses three flags:
* addr: The address of the twitch irc server
  * Default: irc-ws.chat.twitch.tv
* chan: The channel that lurker will be listening to
  * Default: summit1g
* creds: The credentials to connect to your rabbitmq instance
  * Default: user:password@host

Example usage:
```sh
lurker --addr irc-ws.chat.twitch.tv --chan summit1g --creds user:password@host
```



<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/Rwhytsell/Lurker/issues) for a list of proposed features (and known issues).



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Ryan Whytsell - [@Rwhytsell67](https://twitter.com/@Rwhytsell67) - me@ryanw.dev

Project Link: [https://github.com/Rwhytsell/Lurker/](https://github.com/Rwhytsell/Lurker/)


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/RWhytsell/repo.svg?style=for-the-badge
[contributors-url]: https://github.com/Rwhytsell/Lurker/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/RWhytsell/repo.svg?style=for-the-badge
[forks-url]: https://github.com/Rwhytsell/Lurker/network/members
[stars-shield]: https://img.shields.io/github/stars/RWhytsell/repo.svg?style=for-the-badge
[stars-url]: https://github.com/Rwhytsell/Lurker/stargazers
[issues-shield]: https://img.shields.io/github/issues/RWhytsell/repo.svg?style=for-the-badge
[issues-url]: https://github.com/Rwhytsell/Lurker/issues
[license-shield]: https://img.shields.io/github/license/RWhytsell/repo.svg?style=for-the-badge
[license-url]: https://github.com/rwhytsell/Lurker/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/rwhytsell
