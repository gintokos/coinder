import React, { useState } from 'react';
import styles from './links.module.css'

const Links = ({ urls }) => {
  const [activeGroup, setActiveGroup] = useState(null);

  const formatGroupName = (name) => {
    return name
      .split('_')
      .map(word => word.charAt(0).toUpperCase() + word.slice(1))
      .join(' ');
  };

  const handleGroupClick = (groupName) => {
    setActiveGroup(activeGroup === groupName ? null : groupName);
  };

  const getDomain = (url) => {
    try {
      const domain = new URL(url.trim());
      return domain.hostname.replace('www.', '');
    } catch {
      return url.trim();
    }
  };

  const renderLinks = (url) => {
    if (!url) return null;
    
    return url.split(',').map((link, index) => (
      <a 
        key={index}
        href={link.trim()} 
        target="_blank" 
        rel="noopener noreferrer"
        className={styles.link}
      >
        {getDomain(link)}
      </a>
    ));
  };

  return (
    <>
      <h2 className={styles.groupTitle}>Websites</h2>
      {Object.entries(urls).map(([key, value]) => {
        if (!value) return null;
        
        const groupName = formatGroupName(key);
        const isActive = activeGroup === key;

        return (
          <div key={key} className={styles.group}>
            <button
              className={`${styles.button} ${isActive ? styles.activeButton : ''}`}
              onClick={() => handleGroupClick(key)}
            >
              {groupName}
            </button>
            
            {isActive && (
              <div className={styles.linksWrapper}>
                {renderLinks(value)}
              </div>
            )}
          </div>
        );
      })}
    </>
  );
};

export default Links;