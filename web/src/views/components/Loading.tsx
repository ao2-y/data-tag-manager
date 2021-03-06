import React from 'react';

interface OwnProps {
  title?: string;
}

type LoadingProps = OwnProps;

const Loding: React.FC<LoadingProps> = (prop) => {
  return (
    <div className="d-flex">
      <div className="spinner-border" role="status">
        <span className="sr-only">{prop.title || 'Loading...'}</span>
      </div>
      <div className="my-auto ml-2">{prop.title || 'Loading...'}</div>
    </div>
  );
};

export default React.memo(Loding);
