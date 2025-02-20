import { useState, useEffect } from "react";
import "./multiform.css";

const MultiStepForm = () => {
  const [step, setStep] = useState(1);
  const [userId, setUserId] = useState(null);
  const [userData, setUserData] = useState({
    name: "",
    email: "",
    phone: "",
    cpf: "",
    passwordhash: "",
    role: "cliente",
  });

  const [addressData, setAddressData] = useState({
    userId: userId,
    street: "",
    neighborhood: "",
    number: "",
    city: "",
    state: "",
    zipcode: "",
    latitude: "",
    longitude: "",
  });

  const [profileData, setProfileData] = useState({
    userId: userId,
    bio: "",
  });

  const [isFetching, setIsFetching] = useState(false);
  const [formErrors, setFormErrors] = useState([]);

  useEffect(() => {
    if (userId) {
      setAddressData((prev) => ({ ...prev, userId }));
      setProfileData((prev) => ({ ...prev, userId }));
    }
  }, [userId]);

  const sendData = async (url, data) => {
    try {
      const response = await fetch(`http://localhost:4000${url}`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
      });
      const result = await response.json();
      return result;
    } catch (error) {
      console.error("Erro ao enviar dados:", error);
      return { error: "Erro ao enviar dados." };
    }
  };

  const validateStep1 = () => {
    const errors = [];
    if (!userData.name) errors.push("Nome é obrigatório.");
    if (!userData.email) errors.push("Email é obrigatório.");
    if (!userData.phone) errors.push("Telefone é obrigatório.");
    if (!userData.cpf) errors.push("CPF é obrigatório.");
    if (!userData.passwordhash) errors.push("Senha é obrigatória.");
    return errors;
  };

  const validateStep2 = () => {
    const errors = [];
    if (!addressData.zipcode) errors.push("CEP é obrigatório.");
    if (!addressData.number) errors.push("Número é obrigatório.");
    return errors;
  };

  const nextStep = async () => {
    let errors = [];

    if (step === 1) {
      errors = validateStep1();
      if (errors.length === 0) {
        const response = await sendData("/api/users", userData);
        if (response.error) {
          errors.push("Erro ao criar o usuário: " + response.error);
        } else {
          setUserId(response.userId);
        }
      }
    } else if (step === 2) {
      errors = validateStep2();
      if (errors.length === 0) {
        const response = await sendData("/api/addresses", { ...addressData, userId });
        if (response.error) {
          errors.push("Erro ao cadastrar o endereço: " + response.error);
        }
      }
    }

    if (errors.length > 0) {
      setFormErrors(errors);
    } else {
      setFormErrors([]);
      setStep((prev) => prev + 1);
    }
  };

  const prevStep = () => setStep((prev) => prev - 1);

  const handleSubmit = async () => {
    const url = userData.role === "diarista" ? "/api/diarists" : "/api/userprofile";
    const response = await sendData(url, { ...profileData });
    if (response.error) {
      setFormErrors([`Erro ao concluir cadastro: ${response.error}`]);
    } else {
      console.log("Cadastro concluído");
    }
  };

  const fetchAddressByZipcode = async (zipcode) => {
    if (zipcode.length !== 8) return; // Apenas busca se o CEP tiver 8 dígitos
    setIsFetching(true);

    try {
      const response = await fetch(`https://viacep.com.br/ws/${zipcode}/json/`);
      const data = await response.json();

      if (!data.erro) {
        setAddressData((prev) => ({
          ...prev,
          street: data.logradouro || "",
          neighborhood: data.bairro || "",
          city: data.localidade || "",
          state: data.uf || "",
        }));
      } else {
        console.error("CEP não encontrado.");
      }
    } catch (error) {
      console.error("Erro ao buscar endereço:", error);
    } finally {
      setIsFetching(false);
    }
  };

  return (
    <div className="form-container">
      {formErrors.length > 0 && (
        <div className="form-errors">
          <ul>
            {formErrors.map((error, index) => (
              <li key={index}>{error}</li>
            ))}
          </ul>
        </div>
      )}

      {step === 1 && (
        <div className="form-step">
          <h2>Passo 1: Informações do Usuário</h2>
          <input
            type="text"
            placeholder="Nome"
            value={userData.name}
            onChange={(e) => setUserData({ ...userData, name: e.target.value })}
          />
          <input
            type="email"
            placeholder="Email"
            value={userData.email}
            onChange={(e) => setUserData({ ...userData, email: e.target.value })}
          />
          <input
            type="tel"
            placeholder="Telefone"
            value={userData.phone}
            maxLength={11} 
            onChange={(e) => setUserData({ ...userData, phone: e.target.value })}
          />
          <input
            type="text"
            placeholder="CPF"
            value={userData.cpf}
            maxLength={11} 
            onChange={(e) => setUserData({ ...userData, cpf: e.target.value })}
          />
          <input
            type="password"
            placeholder="Senha"
            value={userData.passwordhash}
            onChange={(e) => setUserData({ ...userData, passwordhash: e.target.value })}
          />
          <select
            value={userData.role}
            onChange={(e) => setUserData({ ...userData, role: e.target.value })}
          >
            <option value="cliente">Cliente</option>
            <option value="diarista">Diarista</option>
          </select>
          <button onClick={nextStep}>Próximo</button>
        </div>
      )}

      {step === 2 && (
        <div className="form-step">
          <h2>Passo 2: Endereço</h2>
          <input
            type="text"
            placeholder="CEP"
            value={addressData.zipcode}
            onChange={(e) => {
              const zipcode = e.target.value.replace(/\D/g, ""); // Remove caracteres não numéricos
              setAddressData({ ...addressData, zipcode });
              if (zipcode.length === 8) {
                fetchAddressByZipcode(zipcode);
              }
            }}
          />
          <input
            type="text"
            placeholder="Rua"
            value={addressData.street}
            onChange={(e) => setAddressData({ ...addressData, street: e.target.value })}
            disabled
          />
          <input
            type="text"
            placeholder="Número"
            value={addressData.number}
            onChange={(e) => setAddressData({ ...addressData, number: e.target.value })}
          />
          <input
            type="text"
            placeholder="Bairro"
            value={addressData.neighborhood}
            onChange={(e) => setAddressData({ ...addressData, neighborhood: e.target.value })}
            disabled
          />
          <input
            type="text"
            placeholder="Estado"
            value={addressData.state}
            disabled
          />
          <input
            type="text"
            placeholder="Cidade"
            value={addressData.city}
            disabled
          />
          <button onClick={prevStep}>Voltar</button>
          <button onClick={nextStep}>Próximo</button>
        </div>
      )}

      {step === 3 && (
        <div className="form-step">
          <h2>Passo 3: Perfil do Usuário</h2>
          <textarea
            placeholder="Bio"
            value={profileData.bio}
            onChange={(e) => setProfileData({ ...profileData, bio: e.target.value })}
          ></textarea>
          <textarea
            placeholder="Descrição da Casa"
            value={profileData.housedescription}
            onChange={(e) => setProfileData({ ...profileData, housedescription: e.target.value })}
          ></textarea>
          {userData.role === "diarista" && (
            <>
              <input
                type="number"
                placeholder="Anos de Experiência"
                value={profileData.experienceYears}
                onChange={(e) => setProfileData({ ...profileData, experienceYears: e.target.value })}
              />
                   <input
                type="number"
                placeholder="Preço por Hora"
                value={profileData.pricePerHour}
                onChange={(e) => setProfileData({ ...profileData, pricePerHour: e.target.value })}
              />
            </>
          )}
          <button onClick={prevStep}>Voltar</button>
          <button onClick={handleSubmit}>Finalizar</button>
        </div>
      )}
    </div>
  );
};

export default MultiStepForm;
