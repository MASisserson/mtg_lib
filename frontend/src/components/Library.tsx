import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Card } from '../interfaces/interfaces.js'
import { API_PREFIX } from '../config/api.js';


function CardName(card: Card): React.JSX.Element {
  return (
    <div className="hover:text-blue-500">
      <p className="cursor-pointer">{card.name}</p>
    </div>
  );
}

export default function Library(): React.JSX.Element {
  const [cards, setCards] = useState<Array<Card> | undefined>(undefined);
  const [cardImg, setCardImg] = useState<string | undefined>(undefined)

  useEffect(() => {
    axios.get(API_PREFIX + '/api/library')
      .then(res => {
        setCards(res.data);
        if (res.data && res.data.length > 0) {
          setCardImg(res.data[0].image_uris.normal);
        }
      })
      .catch(err => console.error('Error fetching cards: ', err));
  }, [])

  const hoverHandler = (card: Card) => (e: React.MouseEvent) => {
    e.preventDefault();
    if (e.target) {
      setCardImg(card.image_uris.normal)
    }
  }
  
  return (
    <main className="flex-row">
      <div className="w-1/4">
        <div className="container m-2">
          <img src={cardImg} alt="displayed card" />
        </div>
      </div>
      <ul className="w-3/4">
        {
          cards ?
            cards.map(card => (
              <li key={card.id} onMouseEnter={hoverHandler(card)}>
                <CardName {...card} />
              </li>
            )) :
            <p>Loading...</p>
        }
      </ul>
    </main>
  );
}
