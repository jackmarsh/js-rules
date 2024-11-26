import React from 'react';
import { Button } from 'examples/react/components/button/button';
import { Inline, Stack } from 'examples/react/components/layout/layout';
import jstyles from 'examples/react/styles/styles';

function App() {
  return (
    <div className={jstyles.app_app}>
      My React App
      <div className={jstyles.app_appThingy}/>
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
}

export { App };
