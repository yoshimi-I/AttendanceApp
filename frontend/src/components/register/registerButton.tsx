import React from 'react';
import Button from '@mui/material/Button';

type ButtonType = '作業開始' | '作業終了' | '休憩開始' | '休憩終了';


interface AttendanceButtonProps {
  type: ButtonType;
}

const AttendanceButton: React.FC<AttendanceButtonProps> = ({ type }) => {
  let color: 'primary' | 'secondary' | 'inherit';

  switch (type) {
    case '作業開始':
      color = 'primary';
      break;
    case '作業終了':
      color = 'secondary';
      break;
    default:
      color = 'inherit';
  }

  return (
    <Button
      variant="contained"
      color={color}
      size="large"
      sx={{
        borderRadius: '50%',
        width: '150px',
        height: '150px',
        margin: '2%',
      }}
    >
      {type}
    </Button>
  );
};

export default AttendanceButton;
