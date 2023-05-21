# Life

## Introduction

This is a hobby project focused on exploration of the [Game of Life](https://en.wikipedia.org/wiki/Conway's_Game_of_Life). It is written in Go and utilizes [ebitengine](https://ebitengine.org/) as frontend.

## Configuration

### Config directory structure

| Directory             | Function                                                       |
|:----------------------|:---------------------------------------------------------------|
| `./config/`           | Root config directory. Contains main config file.              |
| `./config/themes/`    | Themes directory. Contains theme files, one file per theme.    |
| `./config/languages/` | Language directory. Contains language files, one per language. |

### Naming convention for config files

| Object           | Rules                                                              | Example             |
|:-----------------|:------------------------------------------------------------------:|:--------------------|
| Main config file | Fixed name                                                         | `config.json`       |
| Language files   | [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) | `en.json`           |
| Theme files      | Unregulated                                                        | `Midnight Sun.json` |

### Main config

Main config file allows the user to tweak most important parameters of the application:

  * world_size - count of cells along one side of the square world. Can be set to any number higher or equal to half of the game window's greater dimension [px] at minimum zoom level (which is the "auto" value for this field). "auto" for window of 720x480 resolution will result in 360x360 map
  * zoom_min - minimum zoom value
  * zoom_max - maximum zoom value
  * language - name of the language file, `.json` extension can be omitted
  * theme - name of the theme file - as above, name alone is enough
  * window - dimensions of the window **w**idth and **h**eight in pixels

Missing and corrupted configuration data is replaced with default values.

### Theme

It is possible to modify a color theme of every interface element. Colors are specified as hexadecimal codes (`#RRGGBB`). An additional value can be appended (`#RRGGBBAA`) to adjust transparency. If standard color code is provided, the game assumes no transparency (`ff`). Hash symbol is optional.

## Game parameters

### Rules

Game of Life is governed by a simple set of rules that determine the state of the cells when world transitions to the next generation. Each cell can assume one of two states: it can either be dead or alive.

The default rules are `23/3`. This notation means that alive cells survive to the next generation if they have 2 or 3 alive neighbours. If they have more or less living neighbours, they die. As for the dead cells, if they have 3 alive neighbours, they become alive. Otherwise, they stay dead. The rules for alive cells are written before the slash and the rules for the dead ones are located after it.

There is a possibility to change rules of Life, using the same pattern (`alive survives / dead becomes alive`). New rules can be submitted into the text entry on control panel. Expression must contain exactly one slash (`/`). Improperly formatted rules will be rejected and the previous ones will reappear in the text entry.

Examples of valid rules:

* `23/3` - survival with 2 or 3 neighbours, rebirth with 3
* `/3` - living cells always die, dead cells can be reborn with 3 alive neighbours
* `23/`- 2 or 3 neighbours assure survival, but no cell can be reborn
* `/`  - rapid mass extinction
* `233/33` - duplicates are removed, so this becomes `23/3`

### Speed

Generation transition rate, expressed as generations per second. It can be adjusted with `+` and `-` buttons. Possible speed values are factors of the game's TPS (Ticks Per Second), which is locked at 60.

### Zoom

Zoom value is the display size of one cell plus its border width in pixels. It can be adjusted with a mouse wheel. Possible zoom values are common divisors of window width and height in range <zoom_min; zoom_max>.

### Jumps

User may choose to fast forward using buttons with roman numerals to jump ahead one (`I`), ten (`X`), fifty (`L`), a hundred (`C`) or a thousand (`M`) generations. Jumps on large worlds may take some time. Fast forward can be cancelled by hitting Play, Reset or Random State buttons.

## License

This software is available under MIT License.
