#include <stdio.h>
#include <stdlib.h>
#include <string.h>


char dec(int enbuffer,int key,int n){
	
	int t = enbuffer;
	char buffer = 0 ;
	for(int i = 1;i<key;i++){
		t *= enbuffer;
		t %= n;
	}
	
	buffer = (char)t;
	
	return buffer;
	
}

char * cma(){
	
	char * debuffer = (char *)malloc(sizeof(char)*200);
	
	return debuffer;
	
}


int main(int argc,char * argv[]){
	int i = 0;
	int k = 0;
	int key ;
	int n;
	FILE *fp,*dfp;
	char filename[20],enfilename[30];
	int enbuffer;
	char * debuffer = cma();
	
	
	if(argv[1] == NULL && argv[2] == NULL){
			printf("no key value");
		if(argv[3] == NULL){
			printf("no file name");
			return -1;
		}
			return -1;
	}
	
	
	key = atoi(argv[1]);
	n = atoi(argv[2]);
	
	strncpy(filename,argv[3],20);
	
	
	printf("key:%d\n  n:%d\n",key,n);
	
	fp = fopen(filename,"r");
	
	while(!feof(fp)){
		fscanf(fp,"%d",&enbuffer);
		debuffer[i] = dec(enbuffer,key,n);
		i++;
	}
	
	
	enfilename[0] = 'd';
	enfilename[1] = 'e';
	
	for(k =2;k<strlen(filename)+2;k++){
		
		enfilename[k] = filename[k-2];
	}
	
	enfilename[k] = 0;
	
	
	printf("filename : %s\n",filename);
	printf("enfilename : %s\n",enfilename);
	
	dfp = fopen(enfilename,"w");
	
	fprintf(dfp,"%s",debuffer);
	
	fclose(dfp);
	
	
	printf("\n"); // 파일로 출력하기만 하면 끝
	fclose(fp);

}