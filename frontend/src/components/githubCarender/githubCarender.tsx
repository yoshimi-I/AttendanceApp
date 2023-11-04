"use client";
import React, { useEffect, useState } from "react";
import Link from "next/link";

import "./styles.css";
import { Box, Container } from "@mui/material";

interface ActivityData {
  day: string;
  activity_time: number;
}

interface ColorCustomization {
  activity0?: string;
  activity1?: string;
  activity2?: string;
  activity3?: string;
  activity4?: string;
}

interface ActivityCalendarProps {
  sampleData?: ActivityData[];
  colorCustomization?: ColorCustomization;
  showMonth?: boolean;
}

export const ActivityCalendar: React.FC<ActivityCalendarProps> = ({
  sampleData,
  colorCustomization,
  showMonth = true,
}) => {
  const [graphData, setGraphData] = useState<number[]>([]);
  const tempGraphData = Array(366).fill(0);
  const [currentYear, setCurrentYear] = useState<number>(2023);
  const [dateText, setDateText] = useState<string[]>([]);
  const [showMonthBar, setShowMonthBar] = useState<boolean>(true);
  const getDayOfYear = (date: string) => {
    var days = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
    var year = parseInt(date.substring(0, 4));
    setCurrentYear(year);
    var month = parseInt(date.substring(5, 7));
    var day = parseInt(date.substring(8));
    if (month > 2 && year % 4 === 0 && (year % 100 !== 0 || year % 400 === 0)) {
      day = day + 1;
    }
    while (month > 0) {
      month = month - 1;
      if (month !== 0) day = day + days[month - 1];
    }
    return day;
  };

  function formatDateForURL(dateString: string | undefined): string {
    if (!dateString) {
      // 必要に応じてデフォルト値やエラーメッセージを返す
      return "/default-value";
    }
    const months: { [key: string]: string } = {
      Jan: "01",
      Feb: "02",
      Mar: "03",
      Apr: "04",
      May: "05",
      Jun: "06",
      Jul: "07",
      Aug: "08",
      Sep: "09",
      Oct: "10",
      Nov: "11",
      Dec: "12",
    };

    const parts = dateString.split(" ");
    const month = months[parts[0]];
    const day = parts[1].replace(",", "").padStart(2, "0");
    const year = parts[2];

    return `/${year}-${month}-${day}`;
  }

  function dateFromDay(day: number) {
    var date = new Date(currentYear, 0);
    let newDate = new Date(date.setDate(day)).toLocaleDateString("en-US", {
      year: "numeric",
      month: "short",
      day: "2-digit",
    });
    return newDate;
  }
  function initialiseDateText() {
    let tempDateTextList = Array(366).fill("");
    for (let i = 1; i <= 365; i++) {
      tempDateTextList[i] = dateFromDay(i);
    }
    setDateText(tempDateTextList);
  }
  const initialise = async () => {
    sampleData?.map((item) => {
      let activityDay = getDayOfYear(item.day);
      tempGraphData[activityDay] = item.activity_time;
    });
    setGraphData(tempGraphData);
  };
  useEffect(() => {
    setShowMonthBar(showMonth);
    initialise();
    // ここで日付テキストを初期化する
    const tempDateTextList = Array.from({ length: 366 }, (_, i) => dateFromDay(i));
    setDateText(tempDateTextList);
  }, [sampleData, showMonth]);
  const matchColorComb = (colorId: number) => {
    if (!colorCustomization) {
      if (colorId >= 4) return "#5105fd";
      else if (colorId == 0) return "#dadada";
      else if (colorId == 2) return "#5105fd69";
      else if (colorId == 1) return "#5105fd52";
      else return "#5105fd99";
    }
    if (colorId >= 4) return colorCustomization.activity4;
    else if (colorId == 0) return colorCustomization.activity0;
    else if (colorId == 1) return colorCustomization.activity1;
    else if (colorId == 2) return colorCustomization.activity2;
    else return colorCustomization.activity3;
  };
  return (
    <Container
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "center",
      }}
    >
      <div className="activity-calender">
        {showMonthBar ? (
          <div className="months-wrapper">
            <span style={{ marginRight: "51px" }}>Jan</span>
            <span style={{ marginRight: "34px" }}>Feb</span>
            <span style={{ marginRight: "33px" }}>Mar</span>
            <span style={{ marginRight: "50px" }}>Apr</span>
            <span style={{ marginRight: "31px" }}>May</span>
            <span style={{ marginRight: "34px" }}>Jun</span>
            <span style={{ marginRight: "53px" }}>Jul</span>
            <span style={{ marginRight: "32px" }}>Aug</span>
            <span style={{ marginRight: "32px" }}>Sep</span>
            <span style={{ marginRight: "50px" }}>Oct</span>
            <span style={{ marginRight: "33px" }}>Nov</span>
            <span>Dec</span>
          </div>
        ) : null}
        <div className="ac-wrapper">
          {[...Array(52)].map((_, i) => {
            return (
              <Box key={i}>
                <div className="aci-wrapper">
                  <Link
                    href={`/mypage/${formatDateForURL(dateText[i * 7 + 1])}`}
                    passHref
                  >
                    <div
                      className={`ac-item day-${i * 7 + 1} ${
                        graphData[i * 7 + 1] !== 0
                          ? `active activity-${
                              graphData[i * 7 + 1] < 4
                                ? graphData[i * 7 + 1]
                                : 4
                            }`
                          : ""
                      }`}
                      style={{
                        background: matchColorComb(graphData[i * 7 + 1]),
                      }}
                    >
                      <span className="tooltiptext">
                        {graphData[i * 7 + 1]} activity on{" "}
                        {dateText.length !== 0 ? dateText[i * 7 + 1] : ""}
                      </span>
                    </div>
                  </Link>
                  <Link
                    href={`/mypage/${formatDateForURL(dateText[i * 7 + 1])}`}
                    passHref
                  >
                    <div
                      className={`ac-item day-${i * 7 + 2} ${
                        graphData[i * 7 + 2] !== 0
                          ? `active activity-${
                              graphData[i * 7 + 2] < 4
                                ? graphData[i * 7 + 2]
                                : 4
                            }`
                          : ""
                      }`}
                      style={{
                        background: matchColorComb(graphData[i * 7 + 2]),
                      }}
                    >
                      <span className="tooltiptext">
                        {graphData[i * 7 + 2]} activity on{" "}
                        {dateText.length !== 0 ? dateText[i * 7 + 2] : ""}
                      </span>
                    </div>
                  </Link>
                  <Link
                    href={`/mypage/${formatDateForURL(dateText[i * 7 + 1])}`}
                    passHref
                  >
                    <div
                      className={`ac-item day-${i * 7 + 3} ${
                        graphData[i * 7 + 3] !== 0
                          ? `active activity-${
                              graphData[i * 7 + 3] < 4
                                ? graphData[i * 7 + 3]
                                : 4
                            }`
                          : ""
                      }`}
                      style={{
                        background: matchColorComb(graphData[i * 7 + 3]),
                      }}
                    >
                      <span className="tooltiptext">
                        {graphData[i * 7 + 3]} activity on{" "}
                        {dateText.length !== 0 ? dateText[i * 7 + 3] : ""}
                      </span>
                    </div>
                  </Link>
                  <Link
                    href={`/mypage/${formatDateForURL(dateText[i * 7 + 1])}`}
                    passHref
                  >
                    <div
                      className={`ac-item day-${i * 7 + 4} ${
                        graphData[i * 7 + 4] !== 0
                          ? `active activity-${
                              graphData[i * 7 + 4] < 4
                                ? graphData[i * 7 + 4]
                                : 4
                            }`
                          : ""
                      }`}
                      style={{
                        background: matchColorComb(graphData[i * 7 + 4]),
                      }}
                    >
                      <span className="tooltiptext">
                        {graphData[i * 7 + 4]} activity on{" "}
                        {dateText.length !== 0 ? dateText[i * 7 + 4] : ""}
                      </span>
                    </div>
                  </Link>
                  <Link
                    href={`/mypage/${formatDateForURL(dateText[i * 7 + 1])}`}
                    passHref
                  >
                    <div
                      className={`ac-item day-${i * 7 + 5} ${
                        graphData[i * 7 + 5] !== 0
                          ? `active activity-${
                              graphData[i * 7 + 5] < 4
                                ? graphData[i * 7 + 5]
                                : 4
                            }`
                          : ""
                      }`}
                      style={{
                        background: matchColorComb(graphData[i * 7 + 5]),
                      }}
                    >
                      <span className="tooltiptext">
                        {graphData[i * 7 + 5]} activity on{" "}
                        {dateText.length !== 0 ? dateText[i * 7 + 5] : ""}
                      </span>
                    </div>
                  </Link>
                  <Link
                    href={`/mypage/${formatDateForURL(dateText[i * 7 + 1])}`}
                    passHref
                  >
                    <div
                      className={`ac-item day-${i * 7 + 6} ${
                        graphData[i * 7 + 6] !== 0
                          ? `active activity-${
                              graphData[i * 7 + 6] < 4
                                ? graphData[i * 7 + 6]
                                : 4
                            }`
                          : ""
                      }`}
                      style={{
                        background: matchColorComb(graphData[i * 7 + 6]),
                      }}
                    >
                      <span className="tooltiptext">
                        {graphData[i * 7 + 6]} activity on{" "}
                        {dateText.length !== 0 ? dateText[i * 7 + 6] : ""}
                      </span>
                    </div>
                  </Link>
                  <Link
                    href={`/mypage/${formatDateForURL(dateText[i * 7 + 1])}`}
                    passHref
                  >
                    <div
                      className={`ac-item day-${i * 7 + 7} ${
                        graphData[i * 7 + 7] !== 0
                          ? `active activity-${
                              graphData[i * 7 + 7] < 4
                                ? graphData[i * 7 + 7]
                                : 4
                            }`
                          : ""
                      }`}
                      style={{
                        background: matchColorComb(graphData[i * 7 + 7]),
                      }}
                    >
                      <span className="tooltiptext">
                        {graphData[i * 7 + 7]} activity on{" "}
                        {dateText.length !== 0 ? dateText[i * 7 + 7] : ""}
                      </span>
                    </div>
                  </Link>
                </div>
              </Box>
            );
          })}
          <div className="aci-wrapper">
            <div
              className={`ac-item day-365 ${
                graphData[365] !== 0
                  ? `active activity-${graphData[365] < 4 ? graphData[365] : 4}`
                  : ""
              }`}
              style={{
                background: matchColorComb(graphData[365]),
              }}
            >
              <span className="tooltiptext">
                {graphData[365]} activity on{" "}
                {dateText.length !== 0 ? dateText[365] : ""}
              </span>
            </div>
          </div>
        </div>
      </div>
    </Container>
  );
};
