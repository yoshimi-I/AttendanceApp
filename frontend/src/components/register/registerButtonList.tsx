import { Box } from '@mui/system';
import AttendanceButton from './registerButton';


function AttendanceButtons() {
  return (
    <Box
      display="flex"

      flexDirection="row"
      gap={2}
      alignItems="center"
      justifyContent="center"
      style={{ marginBottom: '40px' }}
    >
      <AttendanceButton type="作業開始" />
      <AttendanceButton type="作業終了" />
      <AttendanceButton type="休憩開始" />
      <AttendanceButton type="休憩終了" />
    </Box>
  );
}

export default AttendanceButtons;
