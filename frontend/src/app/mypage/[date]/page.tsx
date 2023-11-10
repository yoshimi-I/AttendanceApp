"use client";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";
import ActivitiesByDate from '../../../components/activity/activitiesByDate';
import ReturnButton from '../../../components/return/returnButton';
import { getAuth } from "firebase/auth";

const DateDetail = () => {
  const [data, setData] = useState(null);
  const params = useParams();
  const date = params.date;
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;

  // firebaseのuidを取得
  const auth = getAuth();
  const user = auth.currentUser;
  const userKey = user ? user.uid : null;

  useEffect(() => {
    // URLパラメータが存在する場合にのみAPIを叩く
    if (date) {
      const fetchData = async () => {
        try {
          const response = await fetch(`${baseUrl}/study/history/${userKey}/${date}`);
          if (!response.ok) {
            throw new Error("Network response was not ok");
          }
          const result = await response.json();
          setData(result);
        } catch (error) {
          console.error("Fetch error:", error);
        }
      };

      fetchData();
    }
  }, [date]); // dateが変更されたときにのみ実行

  if (!data) {
    return <div>何もありません</div>;
  }

  return (
    <div>
      <ActivitiesByDate data={data}/>
      <ReturnButton/>
    </div>
  );
};

export default DateDetail;
