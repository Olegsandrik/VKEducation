import React, { useState } from 'react';

import type { Category } from '../../types';

import { useCurrentTime, useProducts } from '../../hooks';
import { Categories, ProductCard } from '../../components';
import { applyCategories, updateCategories } from '../../utils';

import './MainPage.css';

export const MainPage = () => {
    const products = useProducts();
    const currentTime = useCurrentTime();

    const [selectedCategories, setSelectedCategories] = useState<Category[]>(
        []
    );
    const onCategoryClick = (clickedCategory: Category) =>
        setSelectedCategories(
            updateCategories(selectedCategories, clickedCategory)
        );

    return (
        <div className="main-page">
            <h2 className="main-page__title">VK Маркет</h2>
            <h3>{currentTime}</h3>
            <div className="main-page__parameters">
                <Categories
                    selectedCategories={selectedCategories}
                    onCategoryClick={onCategoryClick}
                />
            </div>
            {applyCategories(products, selectedCategories).map((product) => (
                <ProductCard key={product.id} {...product} />
            ))}
        </div>
    );
};
