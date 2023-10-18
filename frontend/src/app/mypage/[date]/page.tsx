'use client'
import { useParams } from 'next/navigation';



const DateDetail = () => {
  const params = useParams()
const date = params.date

  return (
    <div>
      <h1>詳細ページ: {date}</h1>
    </div>
  );
}

export default DateDetail;
