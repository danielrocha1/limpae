import React, { useState } from "react";
import "./register.css"; // Importa o CSS

const PaymentForm = () => {
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    phone: "",
    cpf: "",
    street: "",
    city: "",
    state: "",
    zip: "",
    plan: "basic",
  });

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    fetch("http://localhost:4000/api/checkout", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(formData),
    })
      .then((response) => response.json())
      .then((data) => alert(data.message))
      .catch((error) => console.error(error));
  };

  return (
    <div className="container">
      <h2>Plano de Assinatura</h2>

      <form onSubmit={handleSubmit}>
        {/* Dados Pessoais */}
        <h3>Dados Pessoais</h3>
        <input type="text" name="name" placeholder="Nome Completo" required className="input-field" onChange={handleChange} />
        <input type="email" name="email" placeholder="E-mail" required className="input-field" onChange={handleChange} />
        <input type="text" name="phone" placeholder="Telefone" required className="input-field" onChange={handleChange} />
        <input type="text" name="cpf" placeholder="CPF" required className="input-field" onChange={handleChange} />

        {/* Endereço */}
        <h3>Endereço</h3>
        <input type="text" name="street" placeholder="Rua" required className="input-field" onChange={handleChange} />
        <input type="text" name="city" placeholder="Cidade" required className="input-field" onChange={handleChange} />
        <input type="text" name="state" placeholder="Estado" required className="input-field" onChange={handleChange} />
        <input type="text" name="zip" placeholder="CEP" required className="input-field" onChange={handleChange} />

        {/* Escolha o Plano */}
        <h3>Escolha o Plano</h3>
        <select name="plan" className="input-field" onChange={handleChange}>
          <option value="basic">Básico - R$ 19,90/mês</option>
          <option value="premium">Premium - R$ 49,90/mês</option>
          <option value="free">Gratuito</option>
        </select>

        <button type="submit" className="button-pay">
          Confirmar Assinatura
        </button>
      </form>
    </div>
  );
};

export default PaymentForm;
