import React from 'react';
import {
  Route,
  Routes,
  BrowserRouter as Router,
} from 'react-router';
import { Button } from 'examples/react/components/button';
import { Inline, Stack } from 'examples/react/components/layout';
import jstyles from 'examples/react/styles';

function App() {
  const splash = (
    <div className={jstyles.app_appThingy}>
      <div className={jstyles.card_cardContainer}>
        A card
      </div>
      <Inline gap="2rem">
        <Stack gap="2rem">
          <Button label="Click Me 1" onClick={() => console.log("That felt good!")}/>
          <Button label="Click Me 2" onClick={() => console.log("That felt good!")}/>
        </Stack>
        <Stack gap="2rem">
          <Button label="Click Me 3" onClick={() => console.log("That felt good!")}/>
          <Button label="Click Me 4" onClick={() => console.log("That felt good!")}/>
        </Stack>
      </Inline>
    </div>
  );
  const routes = (
    <Routes>
      <Route path="/" element={splash}/>
      <Route path="/test" element={<div>test</div>}/>
    </Routes>
  );
  return (
    <div className={jstyles.app_app}>
      My React App
      <Router>
        {routes}
      </Router>
    </div>
  );
}

export { App };
