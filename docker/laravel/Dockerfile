FROM php:8.3-cli

WORKDIR /var/www

# Install composer
RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');" && \
    php -r "if (hash_file('sha384', 'composer-setup.php') === 'dac665fdc30fdd8ec78b38b9800061b4150413ff2e3b6f88543c636f7cd84f6db9189d43a81e5503cda447da73c7e5b6') { echo 'Installer verified'; } else { echo 'Installer corrupt'; unlink('composer-setup.php'); } echo PHP_EOL;" && \
    php composer-setup.php && \
    php -r "unlink('composer-setup.php');" && \
    mv composer.phar /usr/local/bin/composer

# Install dependencies
RUN apt-get update && apt-get install -y \
    git \
    zip \
    unzip

# Create laravel project
RUN composer create-project laravel/laravel laravel

EXPOSE 8000

ENTRYPOINT [ "php", "laravel/artisan", "serve" ]

CMD ["--host=0.0.0.0"]