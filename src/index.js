import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';

import Login from './forms/login';
import PaymentPage from './forms/register';

import MultiStepForm from './forms/multiform';
import Dashboard from './dashboard/dashboard'

import MapPage from './diaristmap/map';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    {/* <App /> */}

    {/* <Login /> */}

    {/* <PaymentPage /> */}

    {/* <MultiStepForm /> */}

    <MapPage />

    {/* <Dashboard /> */}
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals

