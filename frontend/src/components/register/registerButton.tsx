import React from 'react';

type ButtonType = '作業開始' | '作業終了' | '休憩開始' | '休憩終了';

interface ActionButtonProps {
  type: ButtonType;
  onClick: (type: ButtonType) => void;
}

const ActionButton: React.FC<ActionButtonProps> = ({ type, onClick }) => {
  return (
    <button onClick={() => onClick(type)} style={{ margin: '10px' }}>
      {type}
    </button>
  );
};

const TaskButtons: React.FC = () => {
  const handleButtonClick = (type: ButtonType) => {
    switch (type) {
      case '作業開始':
        console.log('作業を開始しました');
        // 作業開始のロジックを実装
        break;
      case '作業終了':
        console.log('作業を終了しました');
        // 作業終了のロジックを実装
        break;
      case '休憩開始':
        console.log('休憩を開始しました');
        // 休憩開始のロジックを実装
        break;
      case '休憩終了':
        console.log('休憩を終了しました');
        // 休憩終了のロジックを実装
        break;
    }
  };

  return (
    <div>
      <ActionButton type="作業開始" onClick={handleButtonClick} />
      <ActionButton type="作業終了" onClick={handleButtonClick} />
      <ActionButton type="休憩開始" onClick={handleButtonClick} />
      <ActionButton type="休憩終了" onClick={handleButtonClick} />
    </div>
  );
};

export default TaskButtons;

