"use client";
import { useState } from "react";
import { useRouter } from "next/navigation";
import Avatar from "@mui/material/Avatar";
import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
import Link from "next/link";
import {
  TextField,
  Button,
  Container,
  Typography,
  Box,
  Grid,
  Alert,
} from "@mui/material";
import {
  GoogleAuthProvider,
  getAuth,
  signInWithEmailAndPassword,
  signInWithPopup,
} from "firebase/auth";
import { app } from "../firebase";
import { registerUser } from "../../lib/firebase/getFirebaseInfo";

const LoginPage = () => {
  const router = useRouter();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  const handleLogin = async (e) => {
    e.preventDefault();
    const auth = getAuth();

    try {
      const userCredential = await signInWithEmailAndPassword(
        auth,
        email,
        password
      );
      // ユーザー登録処理を呼び出し、完了を待つ
      await registerUser(userCredential.user);
      // 登録後にマイページへリダイレクト
      router.push("/mypage");
    } catch (error) {
      // エラーが発生した場合はメッセージをstateにセット
      setError(error.message);
    }
  };

  const handleGoogleLogin = async () => {
    const auth = getAuth(app);
    const provider = new GoogleAuthProvider();

    try {
      const userCredential = await signInWithPopup(auth, provider);
      // ユーザー登録処理を呼び出し、完了を待つ
      await registerUser(userCredential.user);
      // 登録後にマイページへリダイレクト
      router.push("/mypage");
    } catch (error) {
      // エラーが発生した場合はメッセージをstateにセット
      console.error("Error during Google sign-in or user registration:", error);
      setError(error.message);
    }
  };

  return (
    <Container component="main" maxWidth="xs">
      <Box
        sx={{
          marginTop: 8,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Avatar sx={{ m: 1, bgcolor: "secondary.main" }}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          ログイン
        </Typography>
        <Box component="form" onSubmit={handleLogin} noValidate sx={{ mt: 1 }}>
          {error && <Alert severity="error">{error}</Alert>}
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="email"
            label="メールアドレス"
            name="email"
            autoComplete="email"
            autoFocus
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="パスワード"
            type="password"
            id="password"
            autoComplete="current-password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >
            ログイン
          </Button>
          <Button
            type="button"
            fullWidth
            variant="contained"
            sx={{
              mt: 1,
              mb: 2,
              textTransform: "none",
              backgroundColor: "rgba(255, 9, 9, 0.802)", // 赤色の透明度を50%に設定
              "&:hover": {
                backgroundColor: "rgba(255, 0, 0, 0.9)", // ホバー時は少し濃くする
              },
            }}
            onClick={handleGoogleLogin}
          >
            Googleアカウントでログイン
          </Button>
          <Grid container>
            <Grid item xs></Grid>
            <Grid item>
              <Link href="/signup" passHref>
                <Typography variant="body2" sx={{ cursor: "pointer" }}>
                  アカウントをお持ちでない方はこちら
                </Typography>
              </Link>
            </Grid>
          </Grid>
        </Box>
      </Box>
    </Container>
  );
};

export default LoginPage;
