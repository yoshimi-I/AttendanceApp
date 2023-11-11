"use client";
import { useRouter } from "next/navigation";
import { Box, Button } from "@mui/material";
import ArrowBackIcon from "@mui/icons-material/ArrowBack";

const ReturnButton = () => {
  const router = useRouter();

  const handleBack = () => {
    router.push("/mypage");
  };

  return (
    <Box display="flex" justifyContent="center" mt={4}>
      <Button
        startIcon={<ArrowBackIcon />}
        onClick={handleBack}
        variant="contained"
        color="primary"
        sx={{
          borderRadius: 20, // ボタンの角をより丸くする
          padding: "10px 20px", // ボタンの内側の余白を調整する
          fontSize: "1rem", // フォントサイズを大きくする
          textTransform: "none", // ボタンのテキストの大文字変換を無効にする
        }}
      >
        戻る
      </Button>
    </Box>
  );
};

export default ReturnButton;
