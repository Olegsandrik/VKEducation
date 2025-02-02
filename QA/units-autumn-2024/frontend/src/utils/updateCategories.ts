import type { Category } from '../types';

export const updateCategories = (
    currentCategories: Category[],
    changedCategories: Category
) => {
    if (currentCategories.includes(changedCategories)) {
        return currentCategories.filter(
            (filter) => filter !== changedCategories
        );
    }

    return [...currentCategories, changedCategories];
};
