"use client";
import { useEffect } from "react";
import { useRouter } from"next/navigation";

const Home = () => {
    const router = useRouter();

    useEffect(() => {
        router.push('/mypage');
    }, [router]);

    return null; // またはローディングコンポーネントを表示
};

export default Home;
