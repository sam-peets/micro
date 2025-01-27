import UserStatus from './components/userStatus'
import './global.css'
import styles from './styles.module.css'

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body>
        <div className={`border center ${styles.container}`}>
          <UserStatus/>
          {children}
        </div>
      </body>
    </html>
  )
}