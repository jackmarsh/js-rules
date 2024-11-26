import React from 'react';
// import styles from './Button.module.css';
import styles from 'examples/react/components/button/styles';

const Button = ({ label, onClick, type = 'button' }) => {
  return (
    <button className={styles.Button_button} onClick={onClick} type={type}>
      {label}
    </button>
  );
};

export { Button };
