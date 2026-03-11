import React from 'react';
import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import CryptChart from './components/chart/crypt_chart';

const buttonStyle = {
  display: 'inline-block',
  padding: '12px 24px',
  backgroundColor: '#007bff',
  color: '#fff',
  textDecoration: 'none',
  borderRadius: '5px',
  fontWeight: 'bold',
  marginTop: '20px'
};

function App() {
  return (
    <BrowserRouter>
      {/* Все, что вне Routes, будет отображаться на каждой странице (как шапка) */}
      <div style={{
        backgroundColor: '#f4f7f6',
        minHeight: '100vh',
        fontFamily: 'sans-serif'
      }}>

        <nav style={{ padding: '20px', borderBottom: '1px solid #ccc', backgroundColor: '#fff' }}>
          <Link style={{ marginRight: '10px' }} to="/">Главная</Link>
          <Link to="/graphs">Мониторинг курса</Link>
        </nav>

        <Routes>
          {/* ГЛАВНАЯ СТРАНИЦА */}
          <Route path="/" element={
            <div style={{ textAlign: 'center', padding: '50px' }}>
              <h1 style={{ color: '#2c3e50' }}>Crypto Dashboard v1.0</h1>
              <p style={{ color: '#7f8c8d' }}>Добро пожаловать в систему мониторинга курсов.</p>
              <Link to="/graphs" style={buttonStyle}>
                Посмотреть графики →
              </Link>
            </div>
          } />

          {/* СТРАНИЦА С ГРАФИКАМИ */}
          <Route path="/graphs" element={
            <div style={{ padding: '40px' }}>
              <header style={{ marginBottom: '30px', textAlign: 'center' }}>
                <h1 style={{ color: '#2c3e50' }}>Live Monitoring</h1>
                <Link to="/" style={{ color: '#007bff', textDecoration: 'none' }}>← Назад</Link>
              </header>

              <main style={{
                maxWidth: '1200px',
                margin: '0 auto',
                backgroundColor: '#fff',
                boxShadow: '0 4px 6px rgba(0,0,0,0.1)',
                borderRadius: '12px',
                padding: '20px'
              }}>
                <CryptChart />
              </main>
            </div>
          } />
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
