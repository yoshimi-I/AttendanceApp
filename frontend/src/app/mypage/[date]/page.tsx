"use client";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";
import { useAuthState } from "react-firebase-hooks/auth";
import { auth } from "../../firebase";
import ActivitiesByDate from "../../../components/activity/activitiesByDate";
import ReturnButton from "../../../components/return/returnButton";
import router from "next/router";

const DateDetail = () => {
  const [data, setData] = useState(null);
  const [user, loading, error] = useAuthState(auth);
  const params = useParams();
  const date = params.date;
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;

  useEffect(() => {
    if (loading) return;
    if (error) {
      console.error("Firebase auth error:", error);
      return;
    }
    if (!user) {
      // ユーザーがログインしていない場合、ログインページにリダイレクト
      router.push("/signin");
      return;
    }

    // ユーザーがログインしている場合、APIリクエストを実行
    const fetchData = async () => {
      try {
        const response = await fetch(
          `${baseUrl}/study/history/${user.uid}/${date}`
        );
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        const result = await response.json();
        setData(result);
      } catch (error) {
        console.error("Fetch error:", error);
      }
    };

    if (date) {
      fetchData();
    }
  }, [user, loading, error, date]);

  if (!data) {
    return
  }

  return (
    <div>
      <ActivitiesByDate data={data} />
      <ReturnButton />
    </div>
  );
};

export default DateDetail;
