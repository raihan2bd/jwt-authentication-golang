<a name="readme-top"></a>
<h1 align='center'>Jwt Authentication with Golang</h1>


# 📗 Table of Contents

- [📗 Table of Contents](#-table-of-contents)
- [ jwt-authentication-golang ](#-about-project-)
  - [🛠 Built With ](#-built-with-)
    - [Tech Stack ](#tech-stack-)
    - [Key Features ](#key-features-)
  - [💻 Getting Started ](#-getting-started-)
    - [Prerequisites](#prerequisites)
    - [Setup](#setup)
    - [Install](#install)
    - [Database](#database)
    - [Usage](#usage)
    - [Build](#build)
  - [👥 Authors ](#-authors-)
  - [🔭 Future Features ](#-future-features-)
  - [🤝 Contributing ](#-contributing-)
  - [⭐️ Show your support ](#️-show-your-support-)
  - [🙏 Acknowledgments ](#-acknowledgments-)
  - [📝 License ](#-license-)


# Jwt Authentication with Golang and Gorm <a name="about-project"></a>


## 🛠 Built With <a name="built-with"></a>
### Tech Stack <a name="tech-stack"></a>
<details>
  <summary>Back End</summary>
  <ul>
    <li>Golang</li>
    <li>PostgreSQL</li>
    <li>Gorm</li>
  </ul>
</details>

### Key Features <a name="key-features"></a>

- Users can login
- User can signup
- User can check if they are authenticated or not

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## 💻 Getting Started <a name="getting-started"></a> 

To get a local copy up and running, follow these steps.

### Prerequisites

In order to run this project you need:
- Then Make sure you have installed [Go (golang)](https://go.dev/dl/) version 1.20.4 or the latest stable version.
- Then make sure you have installed [PostgreSQL](https://www.postgresql.org/) on your local machine if you want to use this project locally.
- Then Create a database called `jwt-go-pos`

### Setup

- Clone this repository to your desired folder:

```sh
  cd your-folder
  https://github.com/raihan2bd/jwt-authentication-golang.git
```

- Before running the project please make sure you create a `.env` file to your project root directory and add `DATABASE_URI`, and `JWT_SECRET_KEY` environment variables to the file. For example:
```
DATABASE_URI="host=localhost port=5432 dbname=jwt-go-pos user=postgres password=your password sslmode=disable"
JWT_SECRET="your jwt secret key"
```

### Install

Install this project with:

- Install the required gems with:

```sh
go mod tidy
```

### Usage

- To run the development server, execute the following command:

```sh
go run .
```

### Build

- To build the project for production-ready run the following command:

```sh
go build -o myapp
```

- To run the build file
  
  ```sh
  ./myapp
  ```


## 👥 Author <a name="author"></a>

👤 **Abu Raihan**

- GitHub: [@raihan2bd](https://github.com/raihan2bd)
- Twitter: [@raihan2bd](https://twitter.com/raihan2bd)
- LinkedIn: [raihan2bd](https://linkedin.com/in/raihan2bd)
  
<p align="right">(<a href="#readme-top">back to top</a>)</p>


## 🔭 Future Features <a name="future-features"></a>

- [ ] **I will add more end-points to this project**

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## 🤝 Contributing <a name="contributing"></a>

Contributions, issues, and feature requests are welcome!

Feel free to check the [issues page](https://github.com/raihan2bd/jwt-authentication-golang/issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## ⭐️ Show your support <a name="support"></a>

If you like this project, please leave a ⭐️

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## 🙏 Acknowledgments <a name="acknowledgements"></a>

I would like to thank [Trevor Sawler](https://www.gocode.ca/) Who helped me a lot to learn Golang.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## 📝 License <a name="license"></a>

This project is [MIT](./LICENSE) licensed.

<p align="right">(<a href="#readme-top">back to top</a>)</p>
