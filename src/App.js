import './HomePage.css';
import Logo from './img/logo.png'
import Agenda from './img/agenda.png'
import Avatar from './img/avatar.png'



export default function HomePage() {
  return (
    <div className="container">
      {/* Topo fixo */}
      <header className="header">
        <img className="logo-left" title="Agenda" src={Agenda} alt="Logo Esquerda" />
        <div>
        <h1>LimpaÊ</h1>
        <p className="header-subtext">Sua casa limpa com profissionais de confiança, perto de você.</p>
        </div>
        <img className="logo-right" title="Login" src={Logo} alt="Logo Direita" />
      </header>
      
      {/* Conteúdo rolável */}
      <main className="content">
        <div className="content-wrapper">
          <h2>Encontre diaristas perto de você!</h2>
          <p>Aqui você encontra profissionais qualificados para serviços domésticos e de limpeza.</p>
          
          {/* Lista de diaristas (Exemplo) */}
          <div className="diaristas-list">
            {[1, 2, 3, 4].map((id) => (
              <div key={id} className="diarista-card">
                <div className="diarista-photo">
                  <img src={Avatar} alt={`Diarista ${id}`} />
                </div>
                <div>
                  <h3>Diarista {id}</h3>
                  <p>Disponível para serviços gerais</p>
                </div>
                <button className="profile-button">Ver Perfil</button>
              </div>
            ))}
          </div>
        </div>
      </main>
      
      {/* Rodapé fixo */}
      <footer className="footer">
        <p>&copy; 2025 Diaristas. Todos os direitos reservados.</p>
      </footer>
    </div>
  );
}
