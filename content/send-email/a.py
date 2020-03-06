import email.message, smtplib
def f_send(s_from, s_to, s_subject, s_content):
   msg = email.message.EmailMessage()
   msg['From'] = s_from
   msg['To'] = s_to
   msg['Subject'] = s_subject
   msg.set_content(s_content)
   pword = input('password: ')
   mua = smtplib.SMTP('smtp.gmail.com')
   mua.starttls()
   mua.login(s_from, pword)
   mua.send_message(msg)
