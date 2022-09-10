FROM --platform=linux/x86_64 mysql:8.0.27

ENV TZ Asia/Tokyo
ENV LC_ALL ja_JP.UTF-8


RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 467B942D3A79BD29 \
     # packageのinstall \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
    locales \
    vim \
    # 日本語を入力できるようにする
    && echo "ja_JP.UTF-8 UTF-8" > /etc/locale.gen  \
    && locale-gen ja_JP.UTF-8 \
    # 不要なmysqlの設定ファイルを削除
    && rm -rf /etc/mysql/conf.d \
    # mysqlのログを入れるファイル作成
    && touch /var/log/mysqld.log \
    && chown mysql:mysql /var/log/mysqld.log


COPY ./my.cnf /etc/mysql/
