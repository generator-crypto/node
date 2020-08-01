#!/usr/bin/env bash

# Генерируем весь state с самого нуля
rm -rf ~/.generatord/ # Выпиливаем все для старта с чистого листа
rm -rf ~/.generatorcli/ # Выпиливаем все для старта с чистого листа

# Инициализируем блокчейн
generatord init genesis --chain-id generator

# Генерим стандартные аккаунты в expect скриптах
#./scripts/generate_jack.exp
#./scripts/generate_alice.exp

# Добавляем первый аккаунт в генезис
generatord add-genesis-account $(appcli keys show coingenerator -a) 10000000000000gnrt,10000000000000stake

# Конфигурируем cli
generatorcli config chain-id generator
generatorcli config output json
generatorcli config indent true
generatorcli config trust-node true

# Генерируем стартовую транзакцию и отправляем ее в validators
echo "abcdef123" | generatord gentx --name coingenerator --amount 100000000stake

generatord collect-gentxs