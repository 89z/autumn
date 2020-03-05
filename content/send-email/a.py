import email.message, smtplib
def m_send(m_from, m_to, m_subject, m_content):
   msg = email.message.EmailMessage()
   msg['From'] = m_from
   msg['To'] = m_to
   msg['Subject'] = m_subject
   msg.set_content(m_content)
   pword = input('password: ')
   mua = smtplib.SMTP('smtp.gmail.com')
   mua.starttls()
   mua.login(m_from, pword)
   mua.send_message(msg)
