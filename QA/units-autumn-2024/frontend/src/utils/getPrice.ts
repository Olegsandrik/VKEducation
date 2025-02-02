import type { PriceSymbol } from '../types';

export const getPrice = (value: number, symbol: PriceSymbol = '₽'): string =>
    `${value.toLocaleString('ru-RU')} ${symbol}`;
