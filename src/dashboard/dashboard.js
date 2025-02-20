import React from "react";
import "./dashboard.css";

const Dashboard = () => {
  return (
    <div className="dashboard-wrapper">
      <div className="dashboard-container">
        <aside className="sidebar">
          <h2>Menu</h2>
          <ul>
            <li>Home</li>
            <li>Relatórios</li>
            <li>Configurações</li>
          </ul>
        </aside>
        <main className="main-content">
          <header className="header">
            <h1>Dashboard</h1>
          </header>
          <section className="stats">
            <div className="card">Usuários: 150</div>
            <div className="card">Vendas: R$ 10.000</div>
            <div className="card">Lucro: R$ 3.500</div>
          </section>
        </main>
      </div>
    </div>
  );
};

export default Dashboard;
