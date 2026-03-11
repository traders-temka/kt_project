import React, { useEffect, useState } from 'react';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts';

const CryptChart = () => {
    const [data, setData] = useState([]);
    const [selectedSymbol, setSelectedSymbol] = useState('BTC');
    const [selectedExchange, setSelectedExchange] = useState('Binance');

    useEffect(() => {
        // Функция для загрузки данных
        const loadData = () => {
            fetch('http://localhost:8080/stat')
                .then(res => res.json()) // 1. Сначала парсим JSON
                .then(json => {          // 2. Теперь работаем с массивом
                    // Преобразуем строковую цену в число и разворачиваем массив
                    const processedData = json.map(item => ({
                        ...item,
                        price: parseFloat(item.price)
                    }));

                    // Сохраняем обработанные данные (хронологический порядок)
                    setData(processedData.reverse());
                })
                .catch(err => console.error("Ошибка загрузки:", err));
        };

        // 1. Загружаем сразу при старте
        loadData();

        // 2. Ставим таймер на обновление каждые 5 секунд
        const interval = setInterval(loadData, 5000);

        // 3. Очищаем таймер, если компонент удалится (важно!)
        return () => clearInterval(interval);
    }, []);

    const filteredData = data.filter(item =>
        item.symbol === selectedSymbol && item.source === selectedExchange
    );

    const symbols = [...new Set(data.map(item => item.symbol))];
    const exchanges = [...new Set(data.map(item => item.source))];

    return (
        <div style={{ width: '100%', height: 500, padding: '20px', backgroundColor: '#fff', borderRadius: '8px' }}>
            <h2 style={{ color: '#333' }}>Мониторинг курса</h2>

            {/* Панель фильтров */}
            <div style={{ marginBottom: '20px', display: 'flex', gap: '10px' }}>
                <select value={selectedSymbol} onChange={(e) => setSelectedSymbol(e.target.value)}>
                    {symbols.map(s => <option key={s} value={s}>{s}</option>)}
                </select>

                <select value={selectedExchange} onChange={(e) => setSelectedExchange(e.target.value)}>
                    {exchanges.map(ex => <option key={ex} value={ex}>{ex}</option>)}
                </select>
            </div>

            <ResponsiveContainer width="100%" height="80%">
                <LineChart data={filteredData}>
                    <CartesianGrid strokeDasharray="3 3" vertical={false} />
                    <XAxis
                        dataKey="timedump"
                        tickFormatter={(str) => new Date(str).toLocaleTimeString()}
                    />
                    <YAxis
                        domain={['auto', 'auto']}
                        tickFormatter={(val) => val.toFixed(2)}
                    />
                    <Tooltip
                        labelFormatter={(label) => new Date(label).toLocaleString()}
                        contentStyle={{ borderRadius: '10px' }}
                    />
                    <Line
                        type="monotone"
                        dataKey="price"
                        stroke="#ff0099cb"
                        strokeWidth={3}
                        dot={false}
                        activeDot={{ r: 6 }}
                        animationDuration={300}
                    />
                </LineChart>
            </ResponsiveContainer>
        </div>
    );
};

export default CryptChart;
