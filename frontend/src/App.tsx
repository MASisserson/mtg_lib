import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Card } from './interfaces/interfaces.js'
import './App.css';

function CardCard(card: Card): React.JSX.Element {
  return (
    <div>
      <img src={card.image_uris.normal} alt="Card"></img>
    </div>
  );
}

function App() {
  const [cards, setCards] = useState<Array<Card> | null>(null);

  useEffect(() => {
    axios.get('/api/library')
      .then(res => setCards(res.data))
      .catch(err => console.error('Error fetching cards: ', err));
  }, [])

  return (
    <div className="App">
      <header>
        <h1>Card Library</h1>
      </header>
      <ul>
        {
          cards ?
            cards.map(card => (
              <li key={card.id}>
                <CardCard {...card} />
              </li>
            )) :
            <p>Loading...</p>
        }
      </ul>
    </div>
  );
}

export default App;
