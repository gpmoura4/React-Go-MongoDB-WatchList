import React from "react";
import "./App.css";
import { Container } from "semantic-ui-react";
import WatchList from "./WatchList";
import { createRoot } from 'react-dom';


function App() {
  return (

    <div>
      {/* <p>TESTANDO</p> */}
      <Container>
        <WatchList></WatchList>
      </Container>
    </div>
  )
}
export default App;
