<!-- DO NOT REMOVE - contributor_list:data:start:["Matt-Gleich"]:end -->

# contrihat

🥧 GitHub contribution graph on the Raspberry Pi Sense Hat LED Matrix. Made in ~1 hour.

![build](https://github.com/Matt-Gleich/contrihat/workflows/build/badge.svg)
![test](https://github.com/Matt-Gleich/contrihat/workflows/test/badge.svg)
![lint](https://github.com/Matt-Gleich/contrihat/workflows/lint/badge.svg)
![release](https://github.com/Matt-Gleich/contrihat/workflows/release/badge.svg)

Display your GitHub contributions on the [raspberry pi sense hat](https://www.raspberrypi.org/products/sense-hat/)!

## 🚀 Setup

1. [Create a personal access token](https://github.com/settings/tokens/new) with `read:user` permission under the user section.
2. On the Raspberry Pi store this token in `~/contrihat-config/pat.txt` (you will have to create the folder as well).
3. Clone this repo and `cd` into it on the Raspberry Pi.
4. If you don't already have it already install docker and docker-compose. This can be done easily by running `sh install-docker.sh` in this repo.
5. Run `docker-compose up -d`.

That's it! You now have your GitHub contribution graph on the Sense Hat!

## ⚙️ Configuration

The configuration file is stored in `~/contrihat-config/config.yml` and is optional. At the moment you can only configure the color levels:

```yaml
levels:
  - '#ffb300'
  - '#0022ff'
  - '#00ff08'
  - '#00ffff'
```

## 🙌 Contributing

Before contributing, please read the [CONTRIBUTING.md file](https://github.com/Matt-Gleich/contrihat/blob/master/CONTRIBUTING.md)

<!-- DO NOT REMOVE - contributor_list:start -->

## 👥 Contributors

- **[@Matt-Gleich](https://github.com/Matt-Gleich)**

<!-- DO NOT REMOVE - contributor_list:end -->
