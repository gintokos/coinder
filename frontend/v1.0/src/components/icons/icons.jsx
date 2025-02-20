export const CommentsIcon = (props) => (
  <svg 
    xmlns="http://www.w3.org/2000/svg" 
    viewBox="0 0 32 32"
    fill="none"
    {...props}
  >
    <g clipPath="url(#clip0_901_1004)">
      <path 
        d="M23.875 25C23.875 25 27.937 29 28.937 30C30.547 31.609 31 31 31 30V8C31 7.447 30.553 7 30 7H8C7.447 7 7 7.447 7 8V26C7 26.553 7.447 27 8 27H22M13 15H25M13 19H18M25 4V2C25 1.437 24.604 1 24 1H2C1.447 1 1 1.447 1 2V20C1 20.553 1.447 21 2 21H7" 
        stroke="currentColor" 
        strokeWidth="2" 
        strokeLinecap="round" 
        strokeLinejoin="round"
      />
    </g>
    <defs>
      <clipPath id="clip0_901_1004">
        <rect width="32" height="32" fill="white"/>
      </clipPath>
    </defs>
  </svg>
)

export const HeartIcon = ({ filled = false, ...props }) =>
  !filled ? (
  <svg 
    viewBox="0 -0.28 512.563 512.563"
    xmlns="http://www.w3.org/2000/svg"
    fill={filled ? "currentColor" : "none"}
    stroke="currentColor"
    {...props}
  >
    <g id="_58_Heart" data-name="58 Heart" transform="translate(0.281)">
      <path 
        d="M512,192c0,141.969-256,320-256,320S0,333.969,0,192c-.281-14.156,1.188-30.031,0-48C-5.281,64.656,64.469,0,144,0c45.531,0,85.625,31.031,112,64C282.375,31.031,322.469,0,368,0c79.531,0,149.281,64.656,144,144C510.812,161.969,512.281,177.844,512,192Zm-64-48c3.969-44-35.812-80-80-80-82.906,0-80,64-80,64a32,32,0,0,1-64,0s2.906-64-80-64c-44.188,0-83.969,36-80,80,1.625,18.062-1.031,34.625.438,48H64c2-.469-11,94.281,191.625,239.219a4.106,4.106,0,0,0,.375-.281,3.776,3.776,0,0,0,.375.281C459,286.281,446,191.531,448,192h-.438C449.031,178.625,446.375,162.062,448,144Z" 
        fillRule="evenodd"
      />
    </g>
  </svg>
) : (
  <svg 
    viewBox="0 -0.28 512.563 512.563"
    xmlns="http://www.w3.org/2000/svg"
    fill="currentColor"
    {...props}
  >
    <g id="_58_Heart" data-name="58 Heart" transform="translate(0.281)">
      <path 
        d="M512,192c0,141.969-256,320-256,320S0,333.969,0,192c-.281-14.156,1.188-30.031,0-48C-5.281,64.656,64.469,0,144,0c45.531,0,85.625,31.031,112,64C282.375,31.031,322.469,0,368,0c79.531,0,149.281,64.656,144,144C510.812,161.969,512.281,177.844,512,192Z" 
      />
    </g>
  </svg>
)