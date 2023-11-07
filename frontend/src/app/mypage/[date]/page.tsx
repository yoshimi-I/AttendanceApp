"use client";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";
import ActivitiesByDate from '../../../components/activity/activitiesByDate';
import ReturnButton from '../../../components/return/returnButton';

const DateDetail = () => {
  const [data, setData] = useState(null);
  const params = useParams();
  const date = params.date;

  useEffect(() => {
    // URLパラメータが存在する場合にのみAPIを叩く
    if (date) {
      const fetchData = async () => {
        try {
          const response = await fetch(`http://localhost:8080/study/history/1/${date}`);
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
    return <div></div>;
  }

  return (
    <div>
      <ActivitiesByDate data={data}/>
      <ReturnButton/>
    </div>
  );
};

export default DateDetail;
