"""
Usage:

```sh
# help
python scripts/mkdir.py -h

# execute in root dir
python scripts/mkdir.py -n [the number of chapter]
```

"""

import argparse
import pathlib
from logging import getLogger, StreamHandler, DEBUG

# set logger
logger = getLogger(__name__)
handler = StreamHandler()
handler.setLevel(DEBUG)
logger.setLevel(DEBUG)
logger.addHandler(handler)
logger.propagate = False


def parse_arguments() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Create directories for reading books with the number of chapters."
    )

    # 引数の設定
    parser.add_argument("num", help="the number of chapters in a book")

    # パース
    args = parser.parse_args()

    return args


def create_dirs(chapter_num: int):
    for i in range(1, chapter_num + 1):
        dir_name = f"src/chap{i:02}"
        p = pathlib.Path(f"./{dir_name}")
        if p.exists():
            logger.debug(f"{dir_name} already exists")
            continue

        p.mkdir(parents=True)  # parents=True 中間ディレクトリの作成を可能にする

    # src 配下のディレクトリの出力
    p = pathlib.Path("./src")
    p.glob(
        "**"
    )  # pathlib モジュールの glob はディレクトリのみ出力(ref: https://note.nkmk.me/python-pathlib-iterdir-glob/)


def rm_dirs(chapter_num: int):
    for i in range(1, chapter_num + 1):
        dir_name = f"chap{i:02}"
        p = pathlib.Path(f"./{dir_name}")
        if not p.exists():
            logger.debug(f"{dir_name} doesn't exists")
            continue
        p.rmdir()


def main() -> None:
    args = parse_arguments()

    chapter_number = int(args.num)
    create_dirs(chapter_number)


if __name__ == "__main__":
    main()
