//
// Created by polar on 2020/2/26.
//

#ifndef LEETCODE_OFFER_2_H
#define LEETCODE_OFFER_2_H


class CMyString {
friend std::ostream & operator<<(std::ostream &os, const CMyString &str);
public:
    explicit CMyString(char *pData = nullptr){
        if(pData == nullptr)
        {
            m_pData = new char[1];
            m_pData[0] = '\0';
        }
        else
        {
            int length = strlen(pData);
            m_pData = new char[length + 1];
            strcpy(m_pData, pData);
        }
    };
    CMyString(const CMyString &str){
        int length = strlen(str.m_pData);
        m_pData = new char[length + 1];
        strcpy(m_pData, str.m_pData);
    };
    CMyString &operator=(CMyString &str);
    ~CMyString(){};

private:
    char *m_pData;
};



#endif //LEETCODE_OFFER_2_H
