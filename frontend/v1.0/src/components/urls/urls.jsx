import { Dropdown, Space } from "antd"
import { CaretRightOutlined } from "@ant-design/icons"
import styles from './urls.module.css'

const items = [
    {
      key: '1',
      type: 'group',
      label: 'websites',
      children: [
        {
          key: '1-1',
          label: <a href="https://example1.com" target="_blank" rel="noopener noreferrer">1st menu item</a>,
        },
        {
          key: '1-2',
          label: <a href="https://example2.com" target="_blank" rel="noopener noreferrer">2nd menu item</a>,
        },
      ],
    },
    {
      key: '2',
      label: 'sub menu',
      children: [
        {
          key: '2-1',
          label: <a href="https://example3.com" target="_blank" rel="noopener noreferrer">3rd menu item</a>,
        },
        {
          key: '2-2',
          label: <a href="https://example4.com" target="_blank" rel="noopener noreferrer">4th menu item</a>,
        },
      ],
    },
    {
      key: '3',
      label: 'disabled sub menu',
      disabled: true,
      children: [
        {
          key: '3-1',
          label: <a href="https://example5.com" target="_blank" rel="noopener noreferrer">5d menu item</a>,
        },
        {
          key: '3-2',
          label: <a href="https://example6.com" target="_blank" rel="noopener noreferrer">6th menu item</a>,
        },
      ],
    },
];

export default function Urls({urls}) {
    return (
        <>
        <Dropdown
            menu={{
            items,
            }}
            placement="top"
            >
            <Space className={styles.space} align={"center"}>
                <span>Websites</span>
                <CaretRightOutlined />
            </Space>
        </Dropdown>
        </>
    )
}