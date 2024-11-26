import React from 'react';
import styles from 'styles';

const Stack = ({
  gap, // The gap between elements in the Stack.
  align, // start, end, center, stretch.
  justify, // start, end, center, space-between, space-around,
  fill = false, // If true, fills available parent container.
  style = {},
  children,
}) => {
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
    flex: fill ? 1 : undefined,
    ...style,
  };
    
  return (
    <div className={styles.Stack_stack} style={combinedStyle}>
      {children}
    </div>
  );
};

export { Stack }
