import { useEffect, useState } from 'react';

export const useCurrentTime = () => {
    const [currentTime, setCurrentTime] = useState(
        new Date().toLocaleTimeString('ru-RU')
    );

    useEffect(() => {
        const interval = setInterval(
            () => setCurrentTime(new Date().toLocaleTimeString('ru-RU')),
            1000
        );

        return () => clearInterval(interval);
    });

    return currentTime;
};
