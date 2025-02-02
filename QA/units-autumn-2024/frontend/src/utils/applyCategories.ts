import type { Category, Product } from '../types';

export const applyCategories = (
    products: Product[],
    categories: Category[]
): Product[] => {
    if (!categories.length) return products;

    return products.filter((product) => categories.includes(product.category));
};
