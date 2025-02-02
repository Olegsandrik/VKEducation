import { Category } from './Category';
import { PriceSymbol } from './Symbol';

export type Product = {
    id: number;
    name: string;
    description: string;
    price: number;
    priceSymbol?: PriceSymbol;
    imgUrl?: string;
    category: Category;
};
