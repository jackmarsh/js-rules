import React, { useState, useEffect } from 'react';
import { Stack } from '.';
import styles from 'styles';

const Inline = ({
  gap, // The gap between elements in the Inline.
  align, // start, end, center, stretch.
  justify, // start, end, center, space-between, space-around,
  switchAt, // Width threshold to switch to stack layout.
  style = {},
  children,
}) => {
  const [isStack, setIsStack] = useState(false);

  useEffect(() => {
    const handleResize = () => {
      if (switchAt) {
        setIsStack(window.innerWidth < switchAt);
      }
    };
    handleResize();
    window.addEventListener('resize', handleResize);
    return () => window.removeEventListener('resize', handleResize);
  }, [switchAt]);
  
  const combinedStyle = {
    gap: gap ? gap : undefined,
    alignItems: align === "start" ? "flex-start" :
      align === "end" ? "flex-end" :
      align === "center" ? "center" : "stretch",
    justifyContent: justify ===  "start" ? "flex-start" :
      justify === "end" ? "flex-end" :
      justify === "center" ? "center" :
      justify === "space-between" ? "space-between" :
      justify === "space-around" ? "space-around" : undefined,
    ...style,
  };
  if (isStack) {
    return (
      <Stack style={combinedStyle}>
        {children}
      </Stack>
    );
  }
  return (
    <div className={styles.Inline_inline} style={combinedStyle}>
      {children}
    </div>
  );
};

export { Inline };
