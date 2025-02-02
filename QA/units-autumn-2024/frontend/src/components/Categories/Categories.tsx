import React, { FC } from 'react';
import cn from 'classnames';

import type { Category } from '../../types';

import './Categories.css';

type CategoriesProps = {
    selectedCategories: Category[];
    onCategoryClick?: (category: Category) => void;
};

const allFilters: Category[] = ['Одежда', 'Для дома', 'Электроника'];

export const Categories: FC<CategoriesProps> = ({
    selectedCategories,
    onCategoryClick,
}) => (
    <div className="categories">
        {allFilters.map((filter) => (
            <div
                key={filter}
                onClick={() => onCategoryClick?.(filter)}
                className={cn('categories__badge', {
                    categories__badge_selected:
                        selectedCategories.includes(filter),
                })}
            >
                {filter}
            </div>
        ))}
    </div>
);
