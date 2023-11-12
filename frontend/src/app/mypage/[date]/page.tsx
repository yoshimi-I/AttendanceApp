"use client";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";
import { useAuthState } from "react-firebase-hooks/auth";
import { auth } from "../../../lib/firebase/firebase";
import ActivitiesByDate from "../../../components/activity/activitiesByDate";
import ReturnButton from "../../../components/return/returnButton";
import router from "next/router";
import { User } from "firebase/auth";

const DateDetail = () => {
  const [data, setData] = useState(null);
  const [user, loading, error] = useAuthState(auth);
  const params = useParams();
  const date = params.date;
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;

  const fetchData = async (user: User) => {
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
    fetchData(user);

    if (date) {
      fetchData(user);
    }
  }, [user, loading, error, date]);

  if (!data) {
    return;
  }

  return (
    <div>
      <ActivitiesByDate data={data} user={user} refreshData={fetchData} />
      <ReturnButton />
    </div>
  );
};

export default DateDetail;
