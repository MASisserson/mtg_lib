import { useState, } from 'react';
import Library from './components/Library';
import SearchBar from './components/SearchBar';
import './App.css';

function App() {
  const [library, setLibrary] = useState<boolean>(true);

  return (
    <div>
      <header>
        <h1>Card Library</h1>
      </header>
      <SearchBar setLibrary={setLibrary} />
      {library ? <Library /> : <p>Test</p>}
    </div>
  );
}

export default App;
