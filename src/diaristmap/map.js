import React, { useState, useEffect } from "react";
import { MapContainer, TileLayer, Marker, Popup } from "react-leaflet";
import L from "leaflet";
import "leaflet/dist/leaflet.css";
import "./styles.css";

const defaultIcon = new L.Icon({
  iconUrl: require("leaflet/dist/images/marker-icon.png"),
  iconSize: [25, 41],
  iconAnchor: [12, 41],
  popupAnchor: [1, -34],
  shadowUrl: require("leaflet/dist/images/marker-shadow.png"),
  shadowSize: [41, 41]
});

const userIcon = new L.Icon({
  iconUrl: require("leaflet/dist/images/marker-icon.png"),
  iconSize: [25, 41],
  iconAnchor: [12, 41],
  popupAnchor: [1, -34],
  shadowUrl: require("leaflet/dist/images/marker-shadow.png"),
  shadowSize: [41, 41],
  className: "green-marker"
});

const diaristas = [
  { id: 1, nome: "Maria Silva", avaliacao: 4.8, coordenadas: [-22.9005, -43.2106] },
  { id: 2, nome: "João Souza", avaliacao: 4.6, coordenadas: [-22.9025, -43.2156] },
  { id: 3, nome: "Ana Lima", avaliacao: 4.5, coordenadas: [-22.9055, -43.2206] },
  { id: 4, nome: "Carlos Mendes", avaliacao: 4.4, coordenadas: [-22.9085, -43.2256] },
  { id: 5, nome: "Fernanda Rocha", avaliacao: 4.3, coordenadas: [-22.9105, -43.2306] }
];

export default function MapPage() {
  const [userLocation, setUserLocation] = useState(null);

  useEffect(() => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          setUserLocation([position.coords.latitude, position.coords.longitude]);
        },
        (error) => {
          console.error("Erro ao obter localização: ", error);
        }
      );
    } else {
      console.error("Geolocalização não suportada pelo navegador.");
    }
  }, []);

  return (
    <div className="container-map">
      <div className="map-content">
        <div className="map">
          <MapContainer center={[-22.9035, -43.2096]} zoom={12} className="map-container">
            <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
            {diaristas.map((diarista) => (
              <Marker key={diarista.id} position={diarista.coordenadas} icon={defaultIcon}>
                <Popup>{diarista.nome} - ⭐ {diarista.avaliacao}</Popup>
              </Marker>
            ))}
            {userLocation && (
              <Marker position={userLocation} icon={userIcon}>
                <Popup>Você está aqui</Popup>
              </Marker>
            )}
          </MapContainer>
        </div>
  
    
        {/* Lista de Diaristas */}
        <div className="map-list-container">
  <h2 className="map-list-title">Diaristas Recomendados</h2>
  <div className="map-list">
    {diaristas.map((diarista) => (
      <div key={diarista.id} className="list-item">
        <div className="list-item-header">
          <span className="list-item-name">{diarista.nome}</span>
          <div className="list-item-rating">
            <span>⭐ {diarista.avaliacao}</span>
          </div>
        </div>
        <div className="list-item-location">
            <span>Distância: 1.5 km</span>
            <br/>
            <br/>
            <span>Valor por Hora: </span>          
        </div>
        <div className="list-item-button">
          <button className="contact-button">Entrar em Contato</button>
        </div>
      </div>
    ))}
  </div>
</div>

      </div>
    </div>
  );
}
